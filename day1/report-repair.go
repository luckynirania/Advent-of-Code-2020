package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	var data []int
	var dataHash map[int]bool = make(map[int]bool)
	s := bufio.NewScanner(f)
	for s.Scan() {
		scanned, err := strconv.Atoi(s.Text())
		check(err)
		data = append(data, scanned)
		dataHash[scanned] = true
	}
	err = s.Err()
	check(err)

	fmt.Println(finder2(data))
	fmt.Println(finder3(data, dataHash))
}

func finder2(s []int) int {
	for _, i := range s {
		for _, j := range s {
			if i+j == 2020 {
				return i * j
			}
		}
	}
	return 0
}

func finder3(s []int, d map[int]bool) int {
	for _, i := range s {
		for _, j := range s {
			k := 2020 - (i + j)
			_, found := d[k]
			if found {
				return i * j * k
			}
		}
	}
	return 0
}
