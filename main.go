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
	problems[15] = problem15
	problems[16] = problem16
	problems[17] = problem17
	problems[18] = problem18
	problems[19] = problem19
	problems[20] = problem20
	problems[21] = problem21
	problems[22] = problem22
	problems[23] = problem23
	problems[24] = problem24
	problems[25] = problem25
	problems[26] = problem26
	problems[27] = problem27
	problems[28] = problem28
	problems[29] = problem29
	problems[30] = problem30
	problems[31] = problem31
	problems[32] = problem32
	problems[33] = problem33
	problems[34] = problem34
	problems[35] = problem35
	problems[36] = problem36
	problems[37] = problem37
	problems[38] = problem38
	problems[39] = problem39
	problems[40] = problem40
	problems[41] = problem41
	problems[42] = problem42
	problems[43] = problem43
	problems[67] = problem67
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

