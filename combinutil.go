package main

import (
	"fmt"
)

type Permutation []int

func NewPermutation(numElem int) Permutation {
	array := make([]int, numElem)
	for i := 0; i < numElem; i++ {
		array[i] = i
	}
	return Permutation(array)
}

func (p Permutation) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Permutation) ToString() string {
	return fmt.Sprintf("%v", p)
}

// returns whether any more permutations remain
func (p Permutation) NextPermutation() bool {

	// find the first element out of order
	i := len(p) - 1
	for ; i > 0 ; i-- {
		if p[i] > p[i-1] { break }
	}

	if i == 0 { return false } // no more permutations left

	// find the successor of p[i-1]
	pi := p[i-1]
 	j := len(p) - 1
	for ; j > 0; j-- {
		if p[j] > pi { break }
	}
	
	
	// swap i-1 & j : i.e. increment p[i-1]
	p.Swap(i-1, j)

	// swap all elements from p[i] to end
	for k := len(p) - 1; i < k; i, k = i+1, k-1 {
		p.Swap(i, k)
	}

	return true
}

func PrintAllPermutations(n int) {
	p := NewPermutation(n)
	println(p.ToString())
	for p.NextPermutation() {
		println(p.ToString())
	}
}
