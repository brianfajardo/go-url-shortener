package main

import (
	"fmt"
	"net/http"

	"github.com/brianfajardo/url-shortener/urlshort"
)

func main() {
	mux := defaultMux()

	pathsToUrl := map[string]string{
		"/cat": "https://en.wikipedia.org/wiki/Cat",
		"/dog": "https://en.wikipedia.org/wiki/Dog",
	}

	mapHandler := urlshort.MapHandler(pathsToUrl, mux)

	yaml := `
- path: /bird
  url: https://en.wikipedia.org/wiki/bird
- path: /whale
  url: https://en.wikipedia.org/wiki/whale
`

	yamlHandler, err := urlshort.YamlHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	return mux
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}
