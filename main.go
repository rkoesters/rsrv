package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr       = flag.String("addr", ":5000", "Address to serve from")
	configFile = flag.String("f", "rsrv.ini", "Configuration file to read from")
	serveCwd   = flag.Bool("d", false, "Ignore '-f' and just serve the current directory.")
)

func main() {
	flag.Parse()

	log.Print("Starting server...")

	ch := make(chan map[string]string)

	go ParseConfig(ch)

	for i := range ch {
		mount, ok := i["mount"]
		if !ok {
			log.Fatalf("error: bad mount point: %v", i)
		}

		log.Printf("mounting: %v", i)

		h := http.StripPrefix(mount, getHandler(i))
		http.Handle(mount, h)
	}

	log.Printf("Serving from: %v", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
