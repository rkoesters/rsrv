package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// ParseConfig starts the process of reading the config file.
func ParseConfig(ch chan map[string]string) {
	defer close(ch)

	parseFile(ch, *configFile)
}

// parseFile opens the given file and parses it.
func parseFile(ch chan map[string]string, fname string) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer f.Close()

	parse(ch, f)
}

// parse reads from the given io.Reader and creates config maps
// that it sends on the channel `ch'.
func parse(ch chan map[string]string, r io.Reader) {
	// rChan allows us to use `range' to read the file line by line.
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
			send(ch, m)
			m = make(map[string]string)
			m["mount"] = line[1 : len(line)-1]

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
	send(ch, m)
}

// send checks to see if the map is not nil before sending it.
func send(ch chan map[string]string, m map[string]string) {
	if m != nil {
		ch <- m
	}
}

// ReadChan is a little helper that allows reading a file line by
// line using `for i := range ch'.
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
