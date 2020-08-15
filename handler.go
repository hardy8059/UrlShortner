package UrlShortner

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-yaml/yaml"
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

type pathUrl struct {
	Path string `yaml:"path,omitempty", json:"path,omitempty"`
	Url  string `json:"url,omitempty"`
}

func FileHandler(Urls []byte, fallback http.Handler, typeofFile string) (http.HandlerFunc, error) {
	var pathUrls []pathUrl
	var err error
	if typeofFile == "yaml" {
		err = yaml.Unmarshal(Urls, &pathUrls)
	} else if typeofFile == "json" {
		err = json.Unmarshal(Urls, &pathUrls)
	} else {
		return nil, errors.New("unable to parse the defined type. please give a yaml or json file")
	}
	if err != nil {
		return nil, err
	}
	pathToUrls := make(map[string]string)
	for _, u := range pathUrls {
		pathToUrls[u.Path] = u.Url
	}
	return MapHandler(pathToUrls, fallback), nil
}
