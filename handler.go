package main

import (
	"log"
	"net/http"
)

// This map contains all of the functions that we can use
// to create handlers.
var newHandler = map[string]func(map[string]string) http.Handler{
	"dir":  newDirHandler,
//	"file": newFileHandler,
}

func getHandler(config map[string]string) http.Handler {
	t, ok := config["type"]
	if !ok {
		log.Fatalf("error: missing type: %v", config)
	}

	// Get the function that creates our handler.
	f, ok := newHandler[t]
	if !ok {
		log.Fatalf("error: unknown type: %v", t)
	}

	return f(config)
}
