package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var imagePath = "/files/image.jpg"
var imageUrl = "https://picsum.photos/1200"

func sendInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "Internal server error")
}

func checkForInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		sendInternalServerError(w)
	}
}

func fetchNewImage(w http.ResponseWriter) {
	result, err := http.Get(imageUrl)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Image not found")
	}

	img, err := os.Create(imagePath)
	checkForInternalServerError(err, w)

	_, err = img.ReadFrom(result.Body)
	checkForInternalServerError(err, w)

	err = img.Sync()
	checkForInternalServerError(err, w)
}

func getImageFromCache(w http.ResponseWriter) *os.File {
	img, err := os.Open(imagePath)
	checkForInternalServerError(err, w)

	return img
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	imageInfo, err := os.Stat(imagePath)

	if err != nil && !errors.Is(err, os.ErrNotExist) {
		sendInternalServerError(w)
	}

	if errors.Is(err, os.ErrNotExist) || imageInfo.ModTime().Day() != time.Now().Day() {
		fetchNewImage(w)
	}

	img := getImageFromCache(w)

	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, img)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("/build")))

	http.HandleFunc("/api/daily-image", imageHandler)

	port := "8090"
	log.Printf("Server starting in port %s", port)
	http.ListenAndServe(":"+port, nil)
}
