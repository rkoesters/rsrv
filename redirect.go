package main

import (
	"log"
	"net/http"
	"strconv"
)

func RedirectHandler(config map[string]string) http.Handler {
	url := mustGet(config, "url")
	status := tryGet(config, "status", "302")

	code, err := strconv.Atoi(status)
	if err != nil {
		log.Fatalf("Redirect: error: parsing status: %v", err)
	}

	return http.RedirectHandler(url, code)
}
