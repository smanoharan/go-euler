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
