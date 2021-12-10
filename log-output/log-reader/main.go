package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func latestLogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}

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

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Checking health")
	pongResult, err := http.Get("http://ping-pong-svc/pingpong/count")

	if pongResult == nil || pongResult.StatusCode != 200 || err != nil {
		fmt.Println("Mr Stark, I don't feel so good.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", latestLogHandler)

	http.ListenAndServe(":5010", nil)
}
