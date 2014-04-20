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

func gcd(a,b int) int {
	if a == 0 { return b }
	if b == 0 { return a } 
	if a < b { return gcd(b % a, a) }
	return gcd(a % b, b)
}

func problem57() string {
	hasMoreDigits := func(numer, denom, powOf10 *big.Int) bool {
		ten := NewBig(10)
		for powOf10.Cmp(denom) == -1 {
			powOf10.Mul(powOf10, ten)
		}
		
		return powOf10.Cmp(numer) == -1
	}

	numer, denom, powOf10 := NewBig(3), NewBig(2), NewBig(1)
	count := 0
	for i := 1; i < 1000; i++ {
		
		// n,d = n+d+d, n+d
		numer.Add(numer, denom)
		denom.Add(numer, denom)
		numer, denom = denom, numer
		
		if hasMoreDigits(numer, denom, powOf10) { count++ }
	}
	return itoa(count)
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

	return strconv.Itoa(sum)
}

func problem60() string {
	type IPv struct {
		prime int
		pow10 int 
	}

	type ConcatGroup []IPv

	const max, maxP = 10*1000, 10*1000*1000 // just a guess
	isComposite := BuildPrimeSieve(maxP)
	primes := make([]IPv, 0, max/5)
	for i,j := 2,10; i < max; i++ {
		if i >= j { j *= 10 }
		if !isComposite.Get(i) {
			primes = append(primes, IPv{i,j}) 
		}
	}

	isPrime := func(i int) bool {
		if i < maxP { return !isComposite.Get(i) }
		return IsMillerRabinPrime(i)
	}

	isConcatPrime := func(i, j IPv) bool {
		cji := j.prime * i.pow10 + i.prime
		cij := i.prime * j.pow10 + j.prime
		return isPrime(cij) && isPrime(cji)
	}

	// build prime pairs
	np := len(primes)
	sets := make([]ConcatGroup, 0, np/3)
	for i := 0; i < np; i++ {
		for j := i+1; j < np; j++ {
			pi, pj := primes[i], primes[j]
			if isConcatPrime(pi, pj) {
				cg := make([]IPv, 2)
				cg[0], cg[1] = pi, pj
				sets = append(sets, cg)
			}
		}
	}

	// build triples, quads and quintuples
	for t := 3; t <= 5; t++ {
		ns := len(sets)
		newSets := make([]ConcatGroup, 0, ns/3)
		
		for s := 0; s < ns; s++ {
			ss := sets[s]
			nss := len(ss)
			for p := 0; p < np; p++ {
				pp := primes[p]
				if pp.prime <= ss[nss-1].prime { continue }
				
				found := true
				for i := 0; i < nss; i++ {
					if !isConcatPrime(pp, ss[i]) { 
						found = false
						break
					}
				}

				if found {
					newbase := make([]IPv, nss, nss+1)
					copy(newbase, ss)
					newbase = append(newbase, pp)
					newSets = append(newSets, newbase)
				}
			}
		}

		sets = newSets
	}

	minSum := -1
	for i := 0; i < len(sets); i++ {
		arr := sets[i]
		sumP := 0
		for j := 0; j < len(arr); j++ {
			sumP += arr[j].prime
		}

		if minSum < 0 || sumP < minSum {
			minSum = sumP
		}
	}

	return itoa(minSum)
}


