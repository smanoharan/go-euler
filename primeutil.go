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
