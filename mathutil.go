package main

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
