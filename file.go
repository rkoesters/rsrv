package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

type file string

func File(config map[string]string) http.Handler {
	return file(mustGet(config, "path"))
}

func (fname file) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(string(fname))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(w, f)
	if err != nil {
		log.Printf("error: io.Copy: %v", err)
	}
}
