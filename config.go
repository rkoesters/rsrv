package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func parseConfig(ch chan map[string]string) {
	defer close(ch)

	f, err := os.Open(*config)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer f.Close()

	conf := bufio.NewReader(f)

	var m map[string]string
	for {
		line, err := conf.ReadString('\n')
		if err == io.EOF {
			if m != nil {
				ch <- m
			}
			return
		} else if err != nil {
			log.Fatal(err)
		}

		line = strings.TrimSpace(line)
		switch {
		case len(line) == 0:
			// Empty line, skip it.

		case strings.HasPrefix(line, "#"):
			// Comment, skip it.

		case strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]"):
			// New header, send the current map and create a new one.
			if m != nil {
				ch <- m
			}

			m = make(map[string]string)
			m["mount"] = strings.Trim(line, "[]")

		case strings.Contains(line, "="):
			// Line is a key=value pair.
			pair := strings.SplitN(line, "=", 2)
			pair[0] = strings.TrimSpace(pair[0])
			pair[1] = strings.TrimSpace(pair[1])
			m[pair[0]] = os.ExpandEnv(pair[1])

		default:
			log.Fatalf("Error reading config: %v", line)
		}
	}
}
