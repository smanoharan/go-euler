package main

import (
	"os"
	"strconv"
)

type problem func() string

const MAX_PROB = 500
var problems [MAX_PROB]problem

func setup_problems() {
	// TODO use reflection to populate this automatically
	problems[1] = problem1
	problems[2] = problem2
	problems[3] = problem3
	problems[4] = problem4
	problems[5] = problem5
	problems[6] = problem6
	problems[7] = problem7
	problems[8] = problem8
	problems[9] = problem9
	problems[10] = problem10
	problems[11] = problem11
	problems[12] = problem12
	problems[13] = problem13
	problems[14] = problem14
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

