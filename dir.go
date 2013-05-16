package main

import (
	"net/http"
)

func newDirHandler(config map[string]string) http.Handler {
	return http.FileServer(http.Dir(mustGet(config, "path")))
}
