package main

func problem71() string {
	const max = 1000*1000
	a, b := 2, 7 // 2/7 < 3/7
	for d := max; d > 0; d-- {
		if d == 7 { continue }
		i, t := 3*d / 7, int(int64(d)*int64(a) / int64(b))
		for i > t && gcd(i,d) != 1 { i-- }
		if i > t { 
			a,b = i,d 
			println(a, b)
		}
	}
	return itoa(a)
}

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

func problem76() string {
	const max = 100
	arr := make([][]int, max+1)
	for i := 0; i <= max; i++ {
		arr[i] = make([]int, max+1)
	}

	// there is only one way to write any number (>=2) as a sum of 1s
	arr[1][0] = 1
	arr[1][1] = 1
	for j := 2; j <= max; j++ { arr[1][j] = 1 }

	for i := 2; i <= max; i++ {
		for j := 0; j < i; j++ { arr[i][j] = arr[i-1][j] }
		for j := i; j <= max; j++ {
			arr[i][j] = arr[i-1][j] + arr[i][j-i] 
		}
	}

	return itoa(arr[max][max] - 1) // exclude the single number sum
}
