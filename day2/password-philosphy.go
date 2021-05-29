package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	count1 := 0
	count2 := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		res := strings.Split(s.Text(), " ")
		if validator1(res) {
			count1++
		}
		if validator2(res) {
			count2++
		}
	}
	err = s.Err()
	check(err)

	fmt.Println(count1)
	fmt.Println(count2)
}

func validator1(input []string) bool {
	rule := input[0]

	temp := strings.Split(rule, "-")
	lower, err := strconv.Atoi(temp[0])
	check(err)
	upper, err := strconv.Atoi(temp[1])
	check(err)

	letter := strings.Split(input[1], ":")

	password := strings.Split(input[2], "")
	letter_count := 0
	for _, let := range password {
		if letter[0] == let {
			letter_count++
		}
	}
	if letter_count < lower || letter_count > upper {
		return false
	}
	return true
}

func validator2(input []string) bool {
	rule := input[0]

	temp := strings.Split(rule, "-")
	lower, err := strconv.Atoi(temp[0])
	check(err)
	upper, err := strconv.Atoi(temp[1])
	check(err)

	letter := strings.Split(input[1], ":")

	password := strings.Split(input[2], "")

	count := 0
	if password[lower-1] == letter[0] {
		count++
	}
	if password[upper-1] == letter[0] {
		count++
	}
	if count == 1 {
		return true
	}
	return false
}
