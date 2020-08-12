package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/hardy8059/UrlShortener/handler.go"
)

func defaultMux() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, World!")
}

func simpleServer(){
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main(){
	mux := defaultMux()
}