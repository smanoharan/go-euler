package main

import (
	"strconv"
	"math"
)

// TODO write mathutil_386.s

func Max2i(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func Max3i(i, j, k int) int {
	return Max2i(i, Max2i(j, k))
}

func Max2ul(i, j uint64) uint64 {
	if i < j {
		return j
	}
	return i
}

func Max3ul(i, j, k uint64) uint64 {
	return Max2ul(i, Max2ul(j, k))
}

func itoa(i int) string {
	return strconv.Itoa(i) 
}

func i64toa(i int64) string {
	return strconv.FormatInt(i, 10)
}

func SqrtI(i int) int {
	return int(math.Sqrt(float64(i)))
}

func SqrtU64(i uint64) uint64 {
	return uint64(math.Sqrt(float64(i)))
}

func PowI64(base, exp int64) int64 {
	return int64(math.Pow(float64(base), float64(exp))) 
}
