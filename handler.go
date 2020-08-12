package UrlShortner

import (
	"fmt"
	"net/http"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method of Request: ", r.Method)
		path := r.URL.Path
		if dest, foundDest := pathToUrls[path]; foundDest {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
