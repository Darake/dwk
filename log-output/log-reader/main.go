package main

import (
	"fmt"
	"net/http"
	"os"
)

func latestLogHandler(w http.ResponseWriter, r *http.Request) {
	hash, err := os.ReadFile("/files/log.txt")

	if err != nil {
		fmt.Println(err)
	}

	pongCount, err := os.ReadFile("/files/pongcounter.txt")

	if err != nil {
		fmt.Println(err)
	}

	output := string(hash) + "\n" + "Ping / Pongs: %d"
	fmt.Fprintf(w, output, int(pongCount[0]))
}

func main() {
	http.HandleFunc("/", latestLogHandler)

	http.ListenAndServe(":5010", nil)
}
