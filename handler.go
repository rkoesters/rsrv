package main

import (
	"log"
	"net/http"
)

// This map contains all of the functions that we can use
// to create handlers.
var handlers = map[string]func(map[string]string) http.Handler{
	"dir":   Dir,
	"file":  File,
	"debug": Debug,
}

func getHandler(config map[string]string) http.Handler {
	t := mustGet(config, "type")

	// Get the function that creates our handler.
	f, ok := handlers[t]
	if !ok {
		log.Fatalf("error: unknown type: %v", t)
	}

	return f(config)
}

// mustGet returns the value from the given map that goes with
// the given key and logs a fatal error if the key doesn't exist in
// the map.
func mustGet(m map[string]string, k string) string {
	v, ok := m[k]
	if !ok {
		log.Fatalf("error: missing key: '%v' in %v", k, m)
	}
	return v
}

func tryGet(m map[string]string, k string, dflt string) string {
	v, ok := m[k]
	if ok {
		return v
	} else {
		return dflt
	}
}
