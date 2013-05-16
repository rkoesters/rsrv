package main

import (
	"net/http"
)

type file string

func File(config map[string]string) http.Handler {
	return file(mustGet(config, "path"))
}

func (path file) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, string(path))
}
