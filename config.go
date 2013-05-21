package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func ParseConfig(ch chan map[string]string) {
	defer close(ch)

	parseFile(ch, *configFile)
}

func parseFile(ch chan map[string]string, fname string) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer f.Close()

	parse(ch, f)
}

func parse(ch chan map[string]string, r io.Reader) {
	rChan := ReadChan(r)

	var m map[string]string
	for line := range rChan {
		switch {
		case len(line) == 0:
			// Empty line, skip it.

		case strings.HasPrefix(line, "#"):
			// Comment, skip it.

		case strings.HasPrefix(line, "<"):
			// Include another file.
			parseFile(ch, line[1:])

		case strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]"):
			// New header, send the current map and create a new one.
			if m != nil {
				ch <- m
			}

			m = make(map[string]string)
			m["mount"] = line[1:len(line)-1]

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

	if m != nil {
		ch <- m
	}
}

func ReadChan(r io.Reader) chan string {
	ch := make(chan string)

	go readChan(r, ch)

	return ch
}

func readChan(r io.Reader, ch chan string) {
	defer close(ch)

	rd := bufio.NewReader(r)

	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			return
		} else if err != nil {
			log.Fatal(err)
		}

		ch <- strings.TrimSpace(line)
	}
}
