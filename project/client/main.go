package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, _ := url.Parse("http://localhost:8090")

	http.Handle("/api/", httputil.NewSingleHostReverseProxy(target))
	http.Handle("/", http.FileServer(http.Dir("/build")))

	port := "8089"
	http.ListenAndServe(":"+port, nil)
}
