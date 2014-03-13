package main

import "math"

func TriNum(n int64) int64 {
	return (n*(n+1)) >> 1
}

func PentaNum(n int64) int64 {
	return (n*(3*n - 1)) >> 1
}

func HexNum(n int64) int64 {
	return n*(2*n - 1)
}

func IntSqrt(n int64) (bool,int64) {
	sq := int64(math.Sqrt(float64(n)))
	return (sq*sq == n), sq
}

func IsTriNum(n int64) bool {
	// 2*n = i*(i+1)
	// i*i + i - 2n = 0
	// determinant = b*b - 4ac = 1 + 4*1*2n = 8n+1
	isSq, det := IntSqrt(8*n + 1)
	// solution is -b +/- det / 2a
	// = (-1 +/- det)/2 so -1 +/- det must be even, so det must be odd
	return isSq && ((det & 1) == 1)
}

func IsPentaNum(n int64) bool {
	// 2*n = 3i*i - i
	// 3i*i - i - 2*n = 0
	// determinant = 1 + 4*3*2n = 24n + 1
	isSq, det := IntSqrt(24*n + 1)
	// solution is 1 +/- det / 6
	// = (1 +/- det) / 6 so 6 must divide 1 + det
	return isSq && ((det+1)%6 == 0)
}

func IsHexNum(n int64) bool {
	// n = 2*i*i - i
	// 2*i*i - i - n = 0
	// determinant = 1 + 4*2*n = 8n + 1
	isSq, det := IntSqrt(8*n + 1)
	// solution is 1 +/- det / 4a
	// = 1 + det must be divisible by 4
	return isSq && (((det+1) & 3) == 3)
}


