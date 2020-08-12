package main

import (
	"fmt"
	"github.com/hardy8059/UrlShortner"
	"log"
	"net/http"
)

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func simpleServer() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	mux := defaultMux()
	pathToUrls := map[string]string{
		"/blog":     "https://buddingengineer.com",
		"/personal": "https://hardikmunjal.com",
	}
	log.Fatal(http.ListenAndServe(":9000", UrlShortner.MapHandler(pathToUrls, mux)))
}
