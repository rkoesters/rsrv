package main

import (
	"net/http"
	"strings"
	"time"
)

func Debug(config map[string]string) http.Handler {
	mount := mustGet(config, "mount")

	mux := NewServeMux(mount)

	mux.HandleFunc("/", debugIndex)
	mux.Handle("/config", fileHandler(*configFile))

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
