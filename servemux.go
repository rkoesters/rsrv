package main

import (
	"net/http"
	"path"
)

// ServeMux is a wrapper around http.ServeMux in order to work well
// with http.StripPrefix.
type ServeMux struct {
	mux    *http.ServeMux
	prefix string
}

func NewServeMux(prefix string) *ServeMux {
	m := new(ServeMux)

	m.prefix = prefix
	m.mux = http.NewServeMux()

	return m
}

func (m *ServeMux) Prefix(s string) string {
	return path.Clean(m.prefix + s)
}

func (m *ServeMux) Handle(pattern string, handler http.Handler) {
	m.mux.Handle(m.Prefix(pattern), handler)
}

func (m *ServeMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.mux.HandleFunc(m.Prefix(pattern), handler)
}

func (m *ServeMux) Handler(r *http.Request) (http.Handler, string) {
	r.URL.Path = m.Prefix(r.URL.Path)
	return m.mux.Handler(r)
}

func (m *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = m.Prefix(r.URL.Path)
	m.mux.ServeHTTP(w, r)
}
