package main

func problem112() string {
	const max = 1000 * 1000 * 1000

	isBouncy := func(i int) bool {
		last := i % 10
		for i /= 10; i > 0 && last == i%10; i /= 10 {
		}
		if i == 0 {
			return false
		}

		cur := i % 10
		isDesc := cur < last
		last = cur

		for i /= 10; i > 0; i /= 10 {
			cur = i % 10
			if (cur != last) && ((cur < last) != isDesc) {
				return true
			}
			last = cur
		}

		return false
	}

	bouncyCount := 0
	for i := 100; i < max; i++ {
		if isBouncy(i) {
			bouncyCount++
		}

		if bouncyCount*100 == i*99 {
			return itoa(i)
		}
	}

	panic("No solution found.")
}

func problem116() string {
	const R, G, B, M = 2, 3, 4, 5
	const max = 50
	ways := make([][]int64, M)
	for c := R; c < M; c++ {
		ways[c] = make([]int64, max+1)
		for i := 0; i <= 3; i++ {
			ways[c][i] = 1
		}
	}

	for i := 2; i <= max; i++ {
		m := Min2i(M, i+1)
		for c := R; c < m; c++ {
			ways[c][i] = ways[c][i-1] + ways[c][i-c] // last square is { black or other colour }
		}
	}

	return i64toa(ways[R][max] + ways[G][max] + ways[B][max] - 3)
}

func problem117() string {
	const numTiles, maxTileSize = 50, 4
	ways := make([]int64, numTiles+1)
	ways[0] = 1

	for i := 1; i <= numTiles; i++ {
		m := Min2i(i, maxTileSize)
		for j := 1; j <= m; j++ {
			ways[i] += ways[i-j]
		}
	}

	return i64toa(ways[numTiles])
}

func problem120() string {
	// Let Q(a,n) = (a-1)^n + (a+1)^n mod a*a
	// Then Q(a,n) = 2 if n is even, 2*a*n mod a*a if n is odd.
	rMaxSum := 0
	for a := 3; a <= 1000; a++ {
		aa := a * a
		rMax := 2
		for n := 1; n < aa; n += 2 {
			rMax = Max2i(rMax, 2*a*n%aa)
		}

		rMaxSum += rMax
	}
	return itoa(rMaxSum)
}
