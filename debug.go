package main

import (
	"net/http"
)

func Debug(config map[string]string) http.Handler {
	mount := mustGet(config, "mount")

	mux := NewServeMux(mount)

	mux.Handle("/config", fileHandler(*configFile))

	return mux
}
