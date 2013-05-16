package main

import (
	"log"
	"net/http"
)

func newDirHandler(config map[string]string) http.Handler {
	p, ok := config["path"]
	if !ok {
		log.Fatalf("dir: error: bad path: %v", p)
	}

	return http.FileServer(http.Dir(p))
}
