package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var randomString string

func generateRandomString() {
	rand.Seed(time.Now().UnixNano())
	characters := make([]rune, 32)
	for i := range characters {
		characters[i] = letters[rand.Intn(len(letters))]
	}
	randomString = string(characters)
}

func writeInInterval(file *os.File) {
	for range time.Tick(time.Second * 5) {
		var newLog = time.Now().UTC().String() + " " + randomString
		file.Truncate(0)
		file.Seek(0, 0)

		_, err := file.WriteString(newLog)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	generateRandomString()

	file, err := os.OpenFile("/files/log.txt", os.O_RDONLY|os.O_CREATE|os.O_WRONLY, 0666)

	defer file.Close()

	if err != nil {
		panic(err)
	}

	writeInInterval(file)
}
