package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var validators = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var ValidEyeColor = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func main() {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	var passports []map[string]string

	var passportText string

	s := bufio.NewScanner(f)
	for s.Scan() {
		temp := s.Text()
		passportText = passportText + " " + temp
		passportText = strings.TrimSpace(passportText)
		if s.Text() == "" {
			var passport map[string]string = make(map[string]string)
			KeyPair := strings.Split(passportText, " ")
			for _, thing := range KeyPair {
				KP := strings.Split(thing, ":")
				passport[KP[0]] = KP[1]
			}
			passports = append(passports, passport)
			passportText = ""
		}
	}

	// don't forget to add an empty line then "---" at the end of input file
	count_1 := 0
	count_2 := 0
	for _, passport := range passports {
		// Part 1
		if validate(passport) {
			count_1 += 1
		}
		// Part 2
		if validate2(passport) {
			count_2 += 1
		}
	}
	fmt.Println(count_1, count_2)

}

func validate(passport map[string]string) bool {
	for _, toValidate := range validators {
		_, found := passport[toValidate]
		if !found {
			return false
		}
	}
	return true
}

func validate2(passport map[string]string) bool {
	value, found := passport["byr"]
	if !found {
		return false
	} else {
		byr, err := strconv.Atoi(value)
		if err != nil {
			return false
		} else {
			if byr < 1919 || byr > 2003 {
				return false
			}
		}
	}

	value, found = passport["iyr"]
	if !found {
		return false
	} else {
		byr, err := strconv.Atoi(value)
		if err != nil {
			return false
		} else {
			if byr < 2009 || byr > 2021 {
				return false
			}
		}
	}

	value, found = passport["eyr"]
	if !found {
		return false
	} else {
		byr, err := strconv.Atoi(value)
		if err != nil {
			return false
		} else {
			if byr < 2019 || byr > 2031 {
				return false
			}
		}
	}

	value, found = passport["hgt"]
	if !found {
		return false
	} else {
		a, _ := regexp.Compile("[0-9]+")
		number := a.FindString(value)

		b, _ := regexp.Compile("[a-z]+")
		dim := b.FindString(value)

		if number == "" || dim == "" {
			return false
		}
		num, _ := strconv.Atoi(number)
		if dim == "cm" {
			if num < 149 || num > 194 {
				return false
			}
		} else if dim == "in" {
			if num < 58 || num > 77 {
				return false
			}

		} else {
			return false
		}
	}

	value, found = passport["hcl"]
	if !found {
		return false
	} else {
		if value[0] != '#' {
			return false
		}
		matched, err := regexp.MatchString("#[0-9]+|[a-f]+", value)
		if err != nil {
			return false
		}
		if !matched {
			return false
		}
	}

	value, found = passport["ecl"]
	if !found {
		return false
	} else {
		if !contains(ValidEyeColor, value) {
			return false
		}
	}

	value, found = passport["pid"]
	if !found {
		return false
	} else {
		if len(value) != 9 {
			return false
		}
		matched, err := regexp.MatchString("[0-9]+", value)
		if err != nil {
			return false
		}
		if !matched {
			return false
		}
	}

	return true

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
