package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	rows int
	cols int
)

func main() {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	var grid []string

	s := bufio.NewScanner(f)
	for s.Scan() {
		temp := s.Text()
		grid = append(grid, temp)
	}

	rows = len(grid)
	cols = len(grid[0])

	// Part 1
	fmt.Println(countTrees(grid, 3, 1))

	// Part 2
	a := countTrees(grid, 1, 1)
	b := countTrees(grid, 3, 1)
	c := countTrees(grid, 5, 1)
	d := countTrees(grid, 7, 1)
	e := countTrees(grid, 1, 2)

	fmt.Println(a * b * c * d * e)
}

func countTrees(grid []string, right int, down int) int {
	tree_count := 0

	row_pos := 0
	col_pos := 0

	for row_pos < rows-1 {
		col_pos = (col_pos + right) % cols
		row_pos += down
		// fmt.Println(row_pos, col_pos)
		if grid[row_pos][col_pos] == '#' {
			tree_count += 1
		}
	}

	return tree_count
}
