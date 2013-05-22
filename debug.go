package main

import (
	"log"
	"net/http"
	"strings"
	"io"
	"time"
)

func DebugHandler(config map[string]string) http.Handler {
	mount := mustGet(config, "mount")

	mux := NewServeMux(mount)

	mux.HandleFunc("/", debugIndex)
	mux.Handle("/config/", DebugConfigHandler(config))

	return mux
}

var debugIndexPage = `
<pre>
<a href="config">Configuration File</a>
</pre>
`

func debugIndex(w http.ResponseWriter, r *http.Request) {
	page := strings.NewReader(debugIndexPage)

	http.ServeContent(w, r, "index.html", time.Now(), page)
}

type debugConfigHandler string

func DebugConfigHandler(config map[string]string) http.Handler {
	return debugConfigHandler(mustGet(config, "mount"))
}

func (d debugConfigHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %v", r.URL.Path)

	for _, file := range configFiles {
		log.Printf("Check if %v equals %v", file, r.URL.Path[len(d+"config/"):])
		if file == r.URL.Path[len(d+"config/"):] {
			http.ServeFile(w, r, string(file))
			return
		}
	}

	for _, file := range configFiles {
		io.WriteString(w, file + "\n")
	}
}
