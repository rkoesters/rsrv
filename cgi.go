package main

import (
	"net/http"
	"net/http/cgi"
	"strings"
)

func CgiHandler(config map[string]string) http.Handler {
	h := new(cgi.Handler)

	h.Path = mustGet(config, "path")
	h.Root = mustGet(config, "mount")
	h.Dir = tryGet(config, "dir", "")

	// Get slice of args.
	sep := tryGet(config, "args_sep", " ")
	args := tryGet(config, "args", "")
	h.Args = strings.Split(args, sep)

	// Get slice of environmental variables to set.
	sep = tryGet(config, "env_sep", " ")
	env := tryGet(config, "env", "")
	h.Env = strings.Split(env, sep)

	return h
}
