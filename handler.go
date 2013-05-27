package main

import (
	"log"
	"net/http"
	"sort"
	"strings"
)

// This map contains all of the functions that we can use
// to create handlers.
var handlers = map[string]func(map[string]string) http.Handler{
	"cgi":      CgiHandler,
	"cmd":      CmdHandler,
	"dir":      DirHandler,
	"file":     FileHandler,
	"redirect": RedirectHandler,
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

// tryGet works like mustGet, except it returns a string dflt instead
// of logging a fatal error.
func tryGet(m map[string]string, k string, dflt string) string {
	v, ok := m[k]
	if ok {
		return v
	} else {
		return dflt
	}
}

// getSlice parses a slice from the given map using the key `k'. First,
// it checks if the map contains `k' verbatim, if it does, it
// returns a slice of k's value split by whitespace. If the key doesn't
// exist, it looks for keys in the form `k[*]' and creates a slice by
// sorting the keys and putting the values into a slice.
func getSlice(m map[string]string, k string) []string {
	v, ok := m[k]
	if ok {
		return strings.Fields(v)
	}

	var keys []string
	for key := range m {
		if strings.HasPrefix(key, k+"[") && strings.HasSuffix(key, "]") {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)

	var vals []string
	for _, key := range keys {
		vals = append(vals, m[key])
	}

	if len(vals) != 0 {
		return vals
	} else {
		return nil
	}
}
