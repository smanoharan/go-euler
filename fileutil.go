package main

import (
	"io/ioutil"
	"strings"
	"strconv"
)

// for small ( < 100 MB) files only. Reads entire file into memory.
func ReadAllLines(filepath string) []string {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil { panic(err) } // only way to handle any file I/O here is to notify the user.
	return strings.Split(string(buf), "\n")
}

func ReadGrid(filepath string) [][]int64 {
	lines := ReadAllLines(filepath)
	grid := make([][]int64, len(lines))
	
	for i, line := range lines {
		if len(line) <= 0 { 
			grid = grid[:i]
			break // stop reading at the first empty line
		}
		
		parts := strings.Split(line, " ")
		grid[i] = make([]int64, len(parts))

		for j, part := range parts {
			num, _ := strconv.Atoi(part)
			grid[i][j] = int64(num)
		}
	}

	return grid
}
