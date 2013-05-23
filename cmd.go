package main

import (
	"log"
	"net/http"
	"os/exec"
)

type cmdHandler string

func CmdHandler(config map[string]string) http.Handler {
	return cmdHandler(mustGet(config, "cmd"))
}

func (c cmdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sh", "-c", string(c))

	cmd.Stdout = w

	err := cmd.Run()
	if err != nil {
		log.Print(err)
	}
}
