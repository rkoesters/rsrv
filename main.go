package main

import (
	"flag"
	"log"
)

var (
	addr = flag.String("addr", ":5000", "Address to server from")
	config = flag.String("f", "rsrv.conf", "Configuration file to read from")
)

func main() {
	flag.Parse()

	log.Print("Starting server...")

	ch := make(chan map[string]string)

	go parseConfig(ch)

	for i := range ch {
		log.Print(i)
	}
}
