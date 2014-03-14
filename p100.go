package main

func problem92() string {
	const max = 10*1000*1000 + 1
	visited := NewBitSet(max+1)
	ends89 := NewBitSet(max+1)

	NextLink := func(i int) int {
		next := 0
		for _, d := range toDigits(i) {
			next += d*d
		}
		return next
	}

	count := 0
	for i := 2; i < max; i++ {
		for j := i; true; j = NextLink(j) {
			if j < max && visited.Get(j) {
				if ends89.Get(j) { count++ }
				break;
			}
			if j == 89 || j == 1 {
				if j == 89 {
					if j < max { ends89.Set(j) }
					count++
				}
				if j < max { visited.Set(j) }
				break;
			}
		}
	}

	return itoa(count)
}

func problem97() string {
	const mod = int64(10*1000*1000*1000)
	
	quickExpMod := func(base, exp int) int64 {
		prod := int64(1)
		for mult,mask := int64(base),1; mask <= exp; mult, mask = SafeMultMod(mult, mult, mod), mask << 1 {
			if (exp & mask) > 0 {
				prod = SafeMultMod(prod, mult, mod)
			}
		}
		return prod
	}

	prime := quickExpMod(2, 7830457)
	prime = (28433*prime + 1) % mod
	return i64toa(prime)
}


