package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

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
	checkForError(err)

	var count int
	err = row.Scan(&count)
	checkForError(err)

	fmt.Fprintf(w, "pong %d", count)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	row := DB.QueryRow("SELECT count FROM counts;")
	var count int
	err := row.Scan(&count)
	checkForError(err)

	fmt.Fprint(w, count)
}

func initDB() {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres",
		os.Getenv("POSTGRES_PASSWORD"),
		"ping-pong-db-svc",
		"5432",
		"postgres")

	var err error
	DB, err = sql.Open("postgres", url)
	checkForError(err)

	err = DB.Ping()
	checkForError(err)

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS counts (name VARCHAR(50) UNIQUE NOT NULL, count int NOT NULL);")
	checkForError(err)

	_, err = DB.Exec("INSERT INTO counts (name, count) VALUES ('pong', 0) ON CONFLICT DO NOTHING;")
	checkForError(err)
}

func main() {
	initDB()

	http.HandleFunc("/pingpong/count", countHandler)
	http.HandleFunc("/pingpong", pongHandler)

	http.ListenAndServe(":5011", nil)
}
