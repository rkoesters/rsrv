package main

import (
	"net/http"
)

func newDirHandler(config map[string]string) http.Handler {
	p := mustGet(config, "path")

	return http.FileServer(http.Dir(p))
}
