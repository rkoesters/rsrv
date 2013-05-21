package main

import (
	"net/http"
)

func DirHandler(config map[string]string) http.Handler {
	return http.FileServer(http.Dir(mustGet(config, "path")))
}
