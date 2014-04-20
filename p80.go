package main

func problem74() string {

	fac := make([]int, 10)
	fac[0] = 1
	for i := 1; i < 10; i++ {
		fac[i] = fac[i-1]*i
	}

	next := func(i int) int {
		res := 0
		for j := i; j > 0; j /= 10 {
			res += fac[j % 10]
		}
		return res
	}

	endpt := make(map[int]int)
	endpt[169] = 3
	endpt[363601] = 3
	endpt[1454] = 3
	endpt[871] = 2
	endpt[45361] = 2
	endpt[872] = 2
	endpt[45362] = 2

	const max = 1000*1000
	count := 0
	for i := 1; i < max; i++ {
		j := 0
		for cur := i; true; j++ {
			_, exists := endpt[cur]
			if exists { j += endpt[cur] }

			nx := next(cur)
			if (exists || cur == nx) {
				if (j == 60) { count++ }
				break
			}
			cur = nx
		}
	}

	return itoa(count)
}

