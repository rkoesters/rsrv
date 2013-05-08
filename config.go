package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func parseConfig(ch chan map[string]string) {
	var m map[string]string

	f, err := os.Open(*config)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer f.Close()

	conf := bufio.NewReader(f)

	for {
		line, err := conf.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if err != nil && err == io.EOF {
			break;
		}

		line = strings.TrimSpace(line)
		switch line[0] {
		case "#":
			// This line is a comment, lets skip it.
		case "[":
			// We reached a new header, so lets create a new map.
			m = make(map[string]string)
			m["mount"] = line[1:len(line)-1]
		default:
			
	}
}
