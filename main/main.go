package main

import (
	"fmt"
	"github.com/hardy8059/UrlShortner"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

func runMapHandlerServer(mux *http.ServeMux) {
	pathToUrls := map[string]string{
		"/blog":     "https://buddingengineer.com",
		"/personal": "https://hardikmunjal.com",
	}
	log.Fatal(http.ListenAndServe(":9000", UrlShortner.MapHandler(pathToUrls, mux)))
}

func runFileHandlerServer(mux *http.ServeMux, path string, typeOfFile string) {
	yamlFile, err := ioutil.ReadFile(filepath.Join(path))
	if err != nil {
		return
	}
	YamlHandler, err := UrlShortner.FileHandler(yamlFile, mux, typeOfFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting Server...")
	fmt.Println("Server Started.")
	log.Fatal(http.ListenAndServe(":9000", YamlHandler))
}

func main() {
	pathToFolders := UrlShortner.InitialisePaths()

	//Define Multiplexer
	mux := defaultMux()

	// Run the Default Map Handler
	//runMapHandlerServer(mux)

	//Get Config File
	var filename string
	fmt.Println("Please Enter the name of configuration file: ")
	fmt.Scanf("%s\n", &filename)
	pathToFile := filepath.Join(pathToFolders.PathToUrlsFolder, filename)
	_, err := os.Stat(pathToFile)
	if os.IsNotExist(err) {
		panic("Please recheck the file name and extension entered. This file doesn't exist: \n" + pathToFile)
	}
	extension := filepath.Ext(filename)[1:]

	// Run File Handler
	runFileHandlerServer(mux, pathToFile, extension)
}
