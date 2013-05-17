package main

import (
	"net/http"
)

type fileHandler string

func File(config map[string]string) http.Handler {
	return fileHandler(mustGet(config, "path"))
}

func (f fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, string(f))
}
