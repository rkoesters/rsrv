package main

import (
	"net/http"
)

func Dir(config map[string]string) http.Handler {
	return http.FileServer(http.Dir(mustGet(config, "path")))
}
