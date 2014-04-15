package main

import (
	"sort"
	"math/big"
	"strings"
	"strconv"
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

func problem58() string {
	
	x := 2   
	for p := 3; p*10 >= 4*x + 1; x++ {
		y := 4*x*x - 2*x + 1
		for i := 0; i < 4; i++ {
			if IsMillerRabinPrime(y + 2*i*x) {
				p++
			}
		}
	}

	return itoa(2*x + 1)
}

func problem59() string {

	chars := strings.Split(ReadAllLines("data/p59.txt")[0], ",")
	encoded := make([]byte, len(chars))
	decoded := make([]byte, len(chars)) 
	for i := range(chars) {
		ei, _ := strconv.Atoi(chars[i])
		encoded[i] = byte(ei)
	}

	// valid (i.e. printable) ASCII chars lie in the range 32 to 126
	asciiMin, asciiMax := byte(32), byte(126)
	asciiSpace := byte(32)

	xorDecode := func(a, b, c int) (bool, int) {
		key := [3]byte{ byte(a), byte(b), byte(c) }
		
		sum, spaceCount, lastSpace := 0, 0, 0
		for i := range(encoded) {
			dec := encoded[i] ^ key[i%3]
			
			if (dec < asciiMin || dec > asciiMax) { return false, -1 }
			if (i == 0) && (dec != 32) && (dec != 34) && (dec != 39) && (dec != 40) && (dec < 65 || dec > 122) { return  false, -1 }
			if dec == asciiSpace { 
				spaceCount++
				sinceLastSpace := i - lastSpace
				if sinceLastSpace > 30 { return false, -1 }
				lastSpace = i
			}

			sum += int(dec)
			decoded[i] = dec
		}

		return spaceCount > 25 && len(encoded) - lastSpace <= 30, sum
	}

	count := 0
	aMin, aMax := 97, 122 
	for i := aMin; i <= aMax; i++ {
		for j := aMin; j <= aMax; j++ {
			for k := aMin; k <= aMax; k++ {
				
				if valid, sum := xorDecode(i, j, k); valid {
					count++
					if count != -1 { 
						words := strings.Split(string(decoded), " ") 
						println(i, j, k, sum, decoded[0], words[0], words[1], words[2], words[3]) 
					}
				}
				
			}
		}
	}

	_, sum := xorDecode(103, 111, 100)
	println(string(decoded))


//	max := 0xffffff // 256*256*256 - 1
	//for ijk := uint(0); ijk <= max; ijk++ {
//		i, j, k := ijk << 16
//		print(i,j,k)
//	}

	return strconv.Itoa(sum)
}
