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

	_ "github.com/lib/pq"
)

const imagePath = "/files/image.jpg"
const imageUrl = "https://picsum.photos/1200"

var DB *sql.DB

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

func imageHandler(w http.ResponseWriter, r *http.Request) {
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

func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := DB.Query("SELECT * FROM todos;")
		checkForError(err)

		var todos []string
		var todo string
		for rows.Next() {
			err := rows.Scan(&todo)
			checkForError(err)

			todos = append(todos, todo)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	case http.MethodPost:
		data, _ := ioutil.ReadAll(r.Body)
		newTodo := string(data)

		if len(newTodo) > 140 {
			log.Println("Too long todo provided: " + newTodo)
			http.Error(w, "Provided todo length is too long (140 charactes max)", http.StatusBadRequest)
			return
		}

		DB.Exec("INSERT INTO todos (description) VALUES ($1)", newTodo)
		log.Println("Inserted: " + newTodo)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func initDB() {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres",
		os.Getenv("POSTGRES_PASSWORD"),
		"project-db-svc",
		"5432",
		"postgres")

	var err error
	DB, err = sql.Open("postgres", url)
	checkForError(err)

	err = DB.Ping()
	checkForError(err)

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS todos (description VARCHAR(140) NOT NULL);")
	checkForError(err)
}

func main() {
	initDB()

	http.HandleFunc("/api/daily-image", imageHandler)
	http.HandleFunc("/api/todos", todosHandler)
	http.HandleFunc("/", defaultHandler)

	port := "8090"
	log.Printf("Server starting in port %s", port)
	http.ListenAndServe(":"+port, nil)
}
