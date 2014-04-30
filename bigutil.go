package main

import "math/big"

func NewBig(val int) *big.Int {
	return big.NewInt(int64(val))
}

func NewRat(numer, denom int) *big.Rat {
	return big.NewRat(int64(numer), int64(denom))
}

func bigExp(base, exp int) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(int64(base)), big.NewInt(int64(exp)), big.NewInt(0))
}

func MultBy(b, c *big.Int) {
	b.Mul(b, c)
}

func DigitalSum(b *big.Int) int {
	const ZERO_CHAR = int('0')

	digSum := 0
	for _, d := range b.String() {
		digSum += int(d) - ZERO_CHAR
	}
	return digSum
}

func IsBigPalindrome(b *big.Int) bool {
	digits := b.String()
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}

// find i*j % mod, without overflow
func SafeMultMod(i, j, mod int64) int64 {
	const max32 = int64(0xffffff)
	ij := int64(0)

	if i < max32 && j < max32 {
		ij = i * j
	} else {
		if i > j {
			i, j = j, i
		}

		// if j is even, then i*j == 2*(j/2)*i
		// else, j = 2k+1, so i*j == (2*(j/2)+1)*i == 2*(j/2)*i + i
		ij = SafeMultMod(i, j>>1, mod) << 1

		if (j & 1) == 1 { // j was odd. Need to adjust i*j
			ij = ij + i
		}
	}
	return ij % mod
}
