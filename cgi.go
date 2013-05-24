package main

import (
	"net/http"
	"net/http/cgi"
)

func CgiHandler(config map[string]string) http.Handler {
	h := new(cgi.Handler)

	h.Path = mustGet(config, "path")
	h.Root = mustGet(config, "mount")
	h.Dir = tryGet(config, "dir", "")

	return h
}
