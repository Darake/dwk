package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func pongHandler(w http.ResponseWriter, r *http.Request) {
	row := DB.QueryRow("SELECT count FROM counts;")
	_, err := DB.Exec("UPDATE counts SET count = count + 1 WHERE name='pong';")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var count int
	err = row.Scan(&count)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "pong %d", count)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	if DB == nil {
		fmt.Println("DB not initialised")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	row := DB.QueryRow("SELECT count FROM counts;")
	var count int
	err := row.Scan(&count)
	checkForError(err)

	fmt.Fprint(w, count)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We are in healthhandler")

	err := DB.Ping()
	if err != nil {
		fmt.Println("oh no bad health")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func initDB() {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres",
		os.Getenv("POSTGRES_PASSWORD"),
		"ping-pong-db-svc",
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

			fmt.Println("db connection active")

			fmt.Println("Checking and creating counts table")
			_, err = DB.Exec("CREATE TABLE IF NOT EXISTS counts (name VARCHAR(50) UNIQUE NOT NULL, count int NOT NULL);")
			checkForError(err)

			fmt.Println("Inserting pong counts if not initialised")
			_, err = DB.Exec("INSERT INTO counts (name, count) VALUES ('pong', 0) ON CONFLICT DO NOTHING;")
			checkForError(err)
			break A
		}
	}
}

func main() {
	go initDB()

	http.HandleFunc("/pingpong/count", countHandler)
	http.HandleFunc("/pingpong", pongHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", defaultHandler)

	http.ListenAndServe(":5011", nil)
}
