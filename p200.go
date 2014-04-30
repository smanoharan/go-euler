package main

func problem191() string {
	const max = 30

	f := make([]int64, max+1) // number of prize strings of length i with no late days.
	g := make([]int64, max+1) // number of prize strings of length i.

	f[0], f[1], f[2] = 1, 2, 4
	g[0], g[1], g[2] = 1, 3, 8

	for i := 3; i <= max; i++ {
		for j := 1; j <= 3; j++ {
			f[i] += f[i-j]
			g[i] += g[i-j]
		}
		g[i] += f[i]
	}

	return i64toa(g[max])
}
