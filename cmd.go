package main

import (
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type cmdHandler struct {
	cmd string
	dir string
}

func CmdHandler(config map[string]string) http.Handler {
	c := new(cmdHandler)

	c.cmd = mustGet(config, "cmd")
	c.dir = tryGet(config, "dir", "")

	return c
}

func (c *cmdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := cmdParse(c.cmd, r)
	cmd := exec.Command("sh", "-c", s)

	cmd.Dir = c.dir
	cmd.Stdin = r.Body
	cmd.Stdout = w
	cmd.Stderr = w

	log.Printf("Running cmd: %v", s)
	err := cmd.Run()
	if err != nil {
		log.Printf("error running: '%v' : %v", s, err)
	}
}

// cmdParse parses a string and replaces substrings in the form `%*'
// with the appropriate substitution.
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

// cmdExpand returns the proper substitution for the given string.
func cmdExpand(s string, r *http.Request) string {
	switch s {
	case "%":
		return "%"
	case "f":
		return cmdSanitize(r.URL.Fragment)
	case "h":
		return cmdSanitize(r.URL.Host)
	case "o":
		return cmdSanitize(r.URL.Opaque)
	case "p":
		return cmdSanitize(r.URL.Path)
	case "q":
		return cmdSanitize(r.URL.RawQuery)
	case "s":
		return cmdSanitize(r.URL.Scheme)
	case "u":
		return cmdSanitize(r.URL.String())
	default:
		return "%" + s
	}
}

// cmdSanitize cleans a string so that it can be safely included into
// a shell command.
func cmdSanitize(s string) string {
	s = strings.Replace(s, `\`, `\\`, -1)
	s = strings.Replace(s, `'`, `'\''`, -1)
	return `'` + s + `'`
}
