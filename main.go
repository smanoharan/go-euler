package main

import (
	"os"
	"strconv"
)

type problem func() string

const MAX_PROB = 500
var problems [MAX_PROB]problem

func setup_problems() {
	problems[1] = problem1
	problems[2] = problem2
	problems[3] = problem3
}

func main() {
	setup_problems()

	args := os.Args
	if len(args) != 2 {
		println("Usage: go-euler <problem-number>")
		return
	}
	
	index, err := strconv.Atoi(args[1])
	if err != nil {
		println(args[1], "is not a number")
		return
	} else if (index < 1 || index > MAX_PROB) {
		println("Please choose a problem number between", 1, MAX_PROB)
		return
	}
	
	prob_fn := problems[index]
	if prob_fn == nil {
		println("Problem", index, "is not yet implemented")
		return
	}

	println(prob_fn())		
}

