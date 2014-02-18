package main

import (
	"strconv"
	"math"
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

func problem2() string {
	sum := 0
	max := 4000000
	for prev, cur := 1, 1; cur <= max; prev, cur = cur, prev + cur {
		if (cur & 1) == 0 { // is even
			sum += cur
		}
	}
	return strconv.Itoa(sum)
}

func problem3() string {
	max := uint64(600851475143)
	sievemax := int(math.Sqrt(float64(max))) + 1
	bitset := NewBitSet(sievemax + 1)
	
	largest_prime := 0
	for p := 2; p <= sievemax; p++ {
		if !bitset.Get(p) {

			for pm := 2*p; pm <= sievemax; pm += p {
				bitset.Set(pm)
			}

			if 0 == (max % uint64(p)) {
				largest_prime = p
			}
		} 
	}

	return strconv.Itoa(largest_prime)
}


