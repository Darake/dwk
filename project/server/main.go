package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	nats "github.com/nats-io/nats.go"
)

const imagePath = "/files/image.jpg"
const imageUrl = "https://picsum.photos/1200"

var DB *sql.DB
var NC *nats.Conn

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type NatsMessage struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

func sendNatsMessage(message string) {
	err := NC.Publish("todo", []byte(message))
	if err != nil {
		log.Println(err)
	}
}

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fetchNewImage(w http.ResponseWriter) {
	result, err := http.Get(imageUrl)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Image not found")
	}

	img, err := os.Create(imagePath)
	checkForError(err)

	_, err = img.ReadFrom(result.Body)
	checkForError(err)

	err = img.Sync()
	checkForError(err)
}

func getImageFromCache(w http.ResponseWriter) *os.File {
	img, err := os.Open(imagePath)
	checkForError(err)

	return img
}

func imageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	imageInfo, err := os.Stat(imagePath)

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}

	if errors.Is(err, os.ErrNotExist) || imageInfo.ModTime().Day() != time.Now().Day() {
		fetchNewImage(w)
	}

	img := getImageFromCache(w)

	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, img)
}

func getTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if DB == nil {
		log.Println("Database not ready")
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	rows, err := DB.Query("SELECT id, description FROM todos WHERE completed = 'false';")
	checkForError(err)

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Description)
		checkForError(err)

		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func newTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if DB == nil {
		log.Println("Database not ready")
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	data, _ := ioutil.ReadAll(r.Body)
	newTodo := string(data)

	if len(newTodo) > 140 {
		log.Println("Too long todo provided: " + newTodo)
		http.Error(w, "Provided todo length is too long (140 charactes max)", http.StatusBadRequest)
		return
	}

	DB.Exec("INSERT INTO todos (description) VALUES ($1)", newTodo)
	log.Println("Inserted: " + newTodo)

	sendNatsMessage("New todo created: " + newTodo)
}

func markTodoDone(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if DB == nil {
		log.Println("Database not ready")
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	todoId := ps.ByName("id")

	DB.Exec("UPDATE todos SET completed = 'true' WHERE id = $1 ;", todoId)

	log.Printf("Todo %s completed", todoId)
	w.WriteHeader(http.StatusOK)

	sendNatsMessage("Todo with id " + todoId + " marked done.")
}

func defaultHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func healthHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Checking health")

	err := DB.Ping()
	if err != nil {
		log.Println("Bad health")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Good health")
	w.WriteHeader(http.StatusOK)
}

func initDB() {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres",
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_URL"),
		"5432",
		"postgres")

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	timeoutExceeded := time.After(5 * time.Minute)
A:
	for {
		select {
		case <-timeoutExceeded:
			fmt.Println("db connection failed after 5min timeout")
			break A
		case <-ticker.C:
			var err error
			DB, err = sql.Open("postgres", url)
			if err != nil {
				fmt.Println("db connection failed")
				continue
			}

			err = DB.Ping()
			if err != nil {
				fmt.Println("db ping failed")
				continue
			}

			fmt.Println("db connection active")

			fmt.Println("Checking and creating todos table")
			_, err = DB.Exec("CREATE TABLE IF NOT EXISTS todos (id serial PRIMARY KEY, description VARCHAR(140) NOT NULL, completed boolean DEFAULT 'false');")
			checkForError(err)

			break A
		}
	}
}

func initNATS() {
	var err error
	NC, err = nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	go initDB()
	go initNATS()

	router := httprouter.New()
	router.GET("/api/todos", getTodos)
	router.POST("/api/todos", newTodo)
	router.PUT("/api/todos/:id", markTodoDone)

	router.GET("/api/daily-image", imageHandler)
	router.GET("/health", healthHandler)
	router.GET("/", defaultHandler)

	port := "8090"
	log.Printf("Server starting in port %s", port)
	http.ListenAndServe(":"+port, router)
}
