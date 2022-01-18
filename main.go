package main

import (
	"fmt"
	"net/http"

	"github.com/brianfajardo/url-shortener/handler"
)

func main() {
	mux := defaultMux()

	pathsToUrl := map[string]string{
		"/cat": "https://en.wikipedia.org/wiki/Cat",
		"/dog": "https://en.wikipedia.org/wiki/Dog",
	}

	mapHandler := handler.MapHandler(pathsToUrl, mux)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	return mux
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}
