package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func parseConfig(ch chan map[string]string) {
	f, err := os.Open(*config)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer f.Close()

	conf := bufio.NewReader(f)

	for {
		line, err := conf.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Print(err)
		}
		if err != nil && err == io.EOF {
			break;
		}

	}
}
