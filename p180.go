package main

func problem179() string {
	const max = 10 * 1000 * 1000

	count := 0
	divCount := make([]int16, max) // ignoring self & 1 (so holds actual-2 for all numbers)
	divCount[1] = -1               // is actual-2

	for i := 2; i < max; i++ {
		for j := 2 * i; j < max; j += i {
			divCount[j]++ // i divides j
		}
		if divCount[i] == divCount[i-1] {
			count++
		}
	}

	return itoa(count)
}
