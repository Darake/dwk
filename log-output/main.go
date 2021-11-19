package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var randomString string
var latestOutput string

func generateRandomString() {
	rand.Seed(time.Now().UnixNano())
	characters := make([]rune, 32)
	for i := range characters {
		characters[i] = letters[rand.Intn(len(letters))]
	}
	randomString = string(characters)
}

func printInInterval() {
	for range time.Tick(time.Second * 5) {
		latestOutput = time.Now().UTC().String() + " " + randomString
		fmt.Println(latestOutput)
	}
}

func latestOutputHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, latestOutput)
}

func main() {
	generateRandomString()
	go printInInterval()

	http.HandleFunc("/", latestOutputHandler)

	http.ListenAndServe(":5010", nil)
}
