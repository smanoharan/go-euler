package main

func problem63() string {
	count := 0
	for base := 1; base < 10; base++ {
		val, baseBig := NewBig(base), NewBig(base)
		for exp := 1; exp == len(val.String()); exp++ {
			println("Incl", base, exp, val.String())
			MultBy(val, baseBig)
			count++
		}
	}
	return itoa(count)
}

func problem67() string {
	return i64toa(sumMax(ReadGrid("data/p67.txt")))
}
