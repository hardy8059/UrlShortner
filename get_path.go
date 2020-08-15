package UrlShortner

import (
	"path/filepath"
	"runtime"
)

type Paths struct {
	BasePath         string
	PathToUrlsFolder string
}

func InitialisePaths() Paths {
	_, currentFile, _, _ := runtime.Caller(0)
	globalPaths := Paths{
		BasePath:         filepath.Dir(currentFile),
		PathToUrlsFolder: filepath.Join(filepath.Dir(currentFile), "urls"),
	}
	return globalPaths
}
