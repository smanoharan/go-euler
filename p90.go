package main

func problem81() string {
	lines := ReadAllLines("data/p81.txt")
	size := len(lines)
	if len(lines[size-1]) == 0 { size-- } // in case last line is blank
	last := size - 1
	
	val := make([][]int, size)
	for r := 0; r < size; r++ {
		val[r] = make([]int, size)
		cells := splitByComma(lines[r])
		for c := 0; c < size; c++ {
			val[r][c] = atoi(cells[c])
		}
	}

	for i := last-1; i >= 0; i-- {
		val[last][i] += val[last][i+1] // no choice for the bottom row
		val[i][last] += val[i+1][last] // no choice for the right col
	}

	for r := last-1; r >= 0; r-- {
		for c := last-1; c >= 0; c-- {
			val[r][c] += Min2i(val[r+1][c], val[r][c+1])
		}
	}

	return itoa(val[0][0])
}

func problem87() string {
	const max = 50*1000*1000
	const primemax, numprimes = 7073, 1000 // 7073*7073 > 50mil ; there are less than 1k primes under 7073
	isComposite := BuildPrimeSieve(primemax)
	primes := make([]int, numprimes)
	np := 0
	for i := 0; i < primemax; i++ {
		if !isComposite.Get(i) { 
			primes[np] = i
			np++
		}
	}

	nums := NewBitSet(max)
	count := 0

	const p4max, p3max, p2max = 90, 370, primemax // respective roots of 50mil 
	for p4 := 0; p4 < np; p4++ {
		pp4 := primes[p4]
		if pp4 >= p4max { break }
		pow4 := pp4*pp4*pp4*pp4

		for p3 := 0; p3 < np; p3++ {
			pp3 := primes[p3]
			if pp3 >= p3max { break }
			pow3 := pp3*pp3*pp3

			pow43 := pow4+pow3
			if pow43 >= max { break }
			
			for p2 := 0; p2 < np; p2++ {
				pp2 := primes[p2]
				pow432 := pow43 + (pp2*pp2)
				if pow432 >= max { break }

				if !nums.Get(pow432) {
					nums.Set(pow432)
					count++
				}
			}
		}
	}

	return itoa(count)
}




