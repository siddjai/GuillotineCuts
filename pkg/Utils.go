package pkg

import (
	"bytes"
	"sort"
)

// Time complexity: O(nlogn)
func ToPermutation(perm []uint8) []uint8 {
	cpperm := make([]int, len(perm))
	for i, p := range perm {
		cpperm[i] = int(p)
	}
	sort.Ints(cpperm)
	pair := make(map[int]int)
	for i, cp := range cpperm {
		pair[cp] = i + 1
	}
	res := make([]uint8, len(perm))
	for i, p := range perm {
		res[i] = uint8(pair[int(p)])
	}
	return res
}

func ToString(perm []uint8) string {
	var buffer bytes.Buffer
	for _, p := range perm {
		buffer.WriteByte(p)
	}
	return buffer.String()
}

func InsertPerm(perm []uint8, a, pos uint8) []uint8 {
	newperm := make([]uint8, len(perm)+1)
	for i := uint8(0); i < pos; i++ {
		newperm[i] = perm[i]
	}
	newperm[pos] = a
	for i := pos; i < uint8(len(perm)); i++ {
		newperm[i+1] = perm[i]
	}
	return newperm
}

func Min(a uint8, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}

func Max(a uint8, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}
