package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func latestLogHandler(w http.ResponseWriter, r *http.Request) {
	hash, err := os.ReadFile("/files/log.txt")

	if err != nil {
		fmt.Println(err)
	}

	pongBody, _ := http.Get("http://ping-pong-svc/pingpong/count")
	pongData, _ := io.ReadAll(pongBody.Body)
	pongCount := string(pongData)

	fmt.Println("Env: " + os.Getenv("MESSAGE"))

	output := os.Getenv("MESSAGE") + "\n" + string(hash) + "\n" + "Ping / Pongs: " + pongCount
	fmt.Fprint(w, output)
}

func main() {
	http.HandleFunc("/", latestLogHandler)

	http.ListenAndServe(":5010", nil)
}
