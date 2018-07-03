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
