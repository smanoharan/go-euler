package main

import "math/big"

func prod(low, high int) *big.Int {
	product := big.NewInt(int64(high))
	for i := low; i < high; i++ {
		product.Mul(product, big.NewInt(int64(i)))
	}
	return product
}

// n choose r (r <= c, not checked)
func nCr(n, r int) *big.Int {
	i, j := Min2i(r, n-r), Max2i(r, n-r)
	// nCr =     n! / (r! (n-r)!)
	//     = (i+j)! / (i! j!)
	//     = (i+j) . (i+j-1) ... (i+1) / i!
	num, denom := prod(i+1, i+j), prod(1, i)
	return num.Quo(num, denom)
}

// effectively, find 40 choose 20
func problem15() string {
	return nCr(40,20).String()
}

func digitSum(digitStr string) int {
	ZERO_CHAR := int('0')
	sum := 0
	for _, c := range digitStr {
		sum += int(c) - ZERO_CHAR
	}
	return sum
}

