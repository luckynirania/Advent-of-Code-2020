package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var group []string
	count1, count2 := 0, 0
	s := bufio.NewScanner(f)

	// Don't forget to put an empty line and "---" at the end of input file
	for s.Scan() {
		group = append(group, s.Text())
		if s.Text() == "" {
			t1, t2 := YesAns(group)
			count1 += t1
			count2 += t2
			group = nil
		}
	}

	// Part 1
	fmt.Println(count1, count2)
}

func YesAns(group []string) (int, int) {
	var sets map[string]int = map[string]int{}
	for _, member := range group {
		temp := strings.Split(member, "")
		for _, each := range temp {
			sets[each] += 1
		}
	}
	GroupMems := len(group) - 1
	CountCommon := 0
	for key := range sets {
		if sets[key] == GroupMems {
			CountCommon += 1
		}
	}
	return len(sets), CountCommon
}
