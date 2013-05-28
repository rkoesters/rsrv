package main

import (
	"net/http"
	"net/http/cgi"
)

func CgiHandler(config map[string]string) http.Handler {
	h := new(cgi.Handler)

	h.Path = mustGet(config, "path")
	h.Root = mustGet(config, "mount")
	h.Dir = tryGet(config, "dir", "")
	h.Args = getSlice(config, "args")
	h.Env = getSlice(config, "env")
	h.InheritEnv = getSlice(config, "inherit")

	h.PathLocationHandler = http.DefaultServeMux

	return h
}
