package main

import (
	"strconv"
)

// sum from start to finish (inclusive), only including every step'th number.
// requires: step \neq 0 and start < finish 
// no requirements are checked
func sum(start, step, finish int) int {
	count := (finish - start) / step
	actualFinish := start + (count * step)
	return (start + actualFinish) * (count+1) / 2 // +1 to include start and finish
}

func problem1() string {
	start := 0
	finish := 999
	return strconv.Itoa(sum(start, 3, finish) + sum(start, 5, finish) - sum(start, 15, finish))
}

func main() {
	println(problem1())
}
