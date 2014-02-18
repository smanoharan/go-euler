package main

// an *unchecked* implementation of bitset. Requires in bounds indices.

const NBIT = 64
const NBITMASK = NBIT - 1

type BitSet []uint64

func NewBitSet(numBits int) BitSet {
	return make([]uint64, numBits/NBIT + 1)
}

func (b *BitSet) toMask(i int) uint64 {
	return 1 << (uint(i) & NBITMASK)
}

func (b *BitSet) Set(i int) {
	(*b)[i/NBIT] |= b.toMask(i)
}

func (b *BitSet) Unset(i int) {
	(*b)[i/NBIT] &^= b.toMask(i)
}

func (b *BitSet) Get(i int) bool {
	return ((*b)[i/NBIT] & b.toMask(i)) > 0
}
