package main

import (
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type cmdHandler string

func CmdHandler(config map[string]string) http.Handler {
	return cmdHandler(mustGet(config, "cmd"))
}

func (c cmdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := cmdParse(string(c), r)
	cmd := exec.Command("sh", "-c", s)

	cmd.Stdin = nil
	cmd.Stdout = w
	cmd.Stderr = w

	log.Printf("Running cmd: %v", s)
	err := cmd.Run()
	if err != nil {
		log.Printf("error running: %v : %v", s, err)
	}
}

func cmdParse(s string, r *http.Request) string {
	var n string
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "%" && i+1 < len(s) {
			i++
			n += cmdExpand(s[i:i+1], r)
		} else {
			n += s[i : i+1]
		}
	}
	return n
}

func cmdExpand(s string, r *http.Request) string {
	switch s {
	case "%":
		return "%"
	case "p":
		return cmdSanitize(r.URL.Path)
	default:
		return "%" + s
	}
}

func cmdSanitize(s string) string {
	s = strings.Replace(s, `'`, `\'`, -1)
	return `'` + s + `'`
}
