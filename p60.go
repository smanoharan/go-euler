package main

import (
	"sort"
	"math/big"
)

func problem52() string {
	toKey := func(i int) int {
		digits := sort.IntSlice(toDigits(i))
		digits.Sort()

		key := 0
		for _, d := range digits {
			key = key*10 + d
		}
		return key
	}

	for i := 1; true; i++ {
		key := toKey(i)
		for j := 2; j <= 6; j++ {
			if key != toKey(i*j) { break }
			if j == 6 { return itoa(i) }
		}
	}

	return "Logic Error"
}

func problem53() string {
	const max, limit = 100, 1000000
	var comb [2][max+1]int

	count := 0

	for n := 0; n <= max; n++ {
		ni := (n & 1)
		nj := 1 - ni // row indices

		comb[ni][0] = 1
		for r := 1; r < n; r++ {
			ncr := comb[nj][r] + comb[nj][r-1]
			if ncr > limit { 
				count++
				ncr = limit + 1 // to avoid overflow
			}
			comb[ni][r] = ncr
		}
		comb[ni][n] = 1
	}

	return itoa(count)
}

func problem55() string {

	reverseBig := func(b *big.Int) *big.Int {
		str := b.String()
		lenStr := len(str)
		strRev := make([]rune, lenStr)
		for i, r := range str { strRev[lenStr - 1 - i] = r }
		
		bReversed := new(big.Int)
		bReversed.SetString(string(strRev), 10)
		return bReversed
	}

	count := 0
	const iMax, jMax = 10000, 51

	for i := 1; i < iMax; i++ {
		ib := NewBig(i)
		count++
		for j := 0; j < jMax; j++ {
			ib.Add(ib, reverseBig(ib)) // compute next Lychrel
			if IsBigPalindrome(ib) { 
				count-- 
				break
			}
		}
	}
	return itoa(count)
}

func problem56() string {
	const max = 100
	maxDsum := 0
	for a := 2; a < max; a++ {
		av, aExpB := NewBig(a), NewBig(a) 
		for b := 2; b < max; b++ {
			MultBy(aExpB, av)
			maxDsum = Max2i(maxDsum, DigitalSum(aExpB))
		}
	}

	return itoa(maxDsum)
}

