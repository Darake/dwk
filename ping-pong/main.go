package main

import (
	"fmt"
	"net/http"
	"os"
)

var filePath = "/files/pongcounter.txt"

func checkForError(err error) {
	if err != nil {
		panic(err)
	}
}

func getPongCount() int {
	data, err := os.ReadFile(filePath)
	if err != nil || string(data) == "" {
		return 0
	}

	count := int(data[0])

	return count
}

func incrementPongCount(currentCount int) {
	newCount := currentCount + 1

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	checkForError(err)

	err = file.Truncate(0)
	checkForError(err)

	_, err = file.Seek(0, 0)
	checkForError(err)

	_, err = file.WriteString(string(newCount))
	checkForError(err)
}

func pongHandler(w http.ResponseWriter, r *http.Request) {
	count := getPongCount()
	incrementPongCount(count)

	fmt.Fprintf(w, "pong %d", count)
}

func main() {
	http.HandleFunc("/", pongHandler)

	http.ListenAndServe(":5011", nil)
}
