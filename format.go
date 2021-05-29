package main

import (
	"bufio"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		// s.Text()
	}
}
