package main

func BuildPrimeSieve(sieveMax int) *BitSet {
	sieve := NewBitSet(sieveMax)

	// 0 and 1 are composites
	sieve.Set(0)
	sieve.Set(1)

	// find all other composites
	for p := 2; p < sieveMax; p++ {
		if !sieve.Get(p) {
			for q := 2*p; q < sieveMax; q+=p {
				sieve.Set(q)
			}
		}
	}

	return &sieve
}

type PrimeAction func(p int)

func IteratePrimes(sieveMax int, action PrimeAction) *BitSet {
	sieve := NewBitSet(sieveMax)

	// 0 and 1 are composites
	sieve.Set(0)
	sieve.Set(1)

	// find all other composites
	for p := 2; p < sieveMax; p++ {
		if !sieve.Get(p) {
			action(p)
			for q := 2*p; q < sieveMax; q+=p {
				sieve.Set(q)
			}
		}
	}

	return &sieve
}

// Check if p is prime according to the Miller-Rabin primality test, with a = 2, 7 and 61.
// Miller-Rabin test (with the 3 a-values) is accurate for all 32-bit integers.
func IsMillerRabinPrime(p int) bool {
	if p == 2 { return true }
	if p < 2 || (p & 1) == 0 { return false } // is Even
	pu := uint64(p)

	// find max { s } s.t. 2^s | q
	s,d := uint64(0),pu-1
	for ; d & 1 == 0 ; d >>= 1 {
		s++
	}

	fastExpMod := func(base, exp, mod uint64) uint64 {
		const u0,u1 = uint64(0), uint64(1)

		if exp == u0 { return u1 }

		result := u1
		for b,e := base,u1; e <= exp; b,e = b*b % mod,e<<1 {
			if exp & e != 0 {
				result = (result * b) % mod
			}
		}
		return result
	}

	isPrimeByWitness := func(a uint64) bool {
		if a >= pu { return true }

		aPowD := fastExpMod(a, d, pu)
		if aPowD == 1 { return true }
		
		for r := uint64(0); r < s; r++ {
			if aPowD == pu-1 { return true }
			aPowD = (aPowD * aPowD) % pu
		}
		return false
	}

	return isPrimeByWitness(2) && isPrimeByWitness(7) && isPrimeByWitness(61)
}


func IsMillerRabinPrimeTest() string {
	const max = 50*1000*1000
	expIsComposite := BuildPrimeSieve(max)
	println("Table Built")

	for p := 1; p < max; p++ {
		const step = max / 100
		if (p % step) == 0 {
			print("  >  ", p / step, "\t% [", p, "]\n") 
		}

		actIsComposite := ! IsMillerRabinPrime(p)
		if actIsComposite != expIsComposite.Get(p) {
			println("Failed", p, expIsComposite.Get(p))
			panic(itoa(p))
		}
	}
	return "Done, with no failures"
}


