package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr   = flag.String("addr", ":5000", "Address to server from")
	config = flag.String("f", "rsrv.conf", "Configuration file to read from")
)

func main() {
	flag.Parse()

	log.Print("Starting server...")

	ch := make(chan map[string]string)

	go parseConfig(ch)

	for i := range ch {
		mount, ok := i["mount"]
		if !ok {
			log.Fatalf("error: bad mount point: %v", i)
		}

		h := http.StripPrefix(mount, getHandler(i))
		http.Handle(mount, h)
	}

	log.Printf("Serving from: %v", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
