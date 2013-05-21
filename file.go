package main

import (
	"net/http"
)

type fileHandler string

func FileHandler(config map[string]string) http.Handler {
	return fileHandler(mustGet(config, "path"))
}

func (f fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, string(f))
}
