package eqv

import "bytes"

func complement(perm []uint8) []uint8 {
	n := uint8(len(perm))
	permC := make([]uint8, n)
	for i := uint8(0); i < n; i++ {
		permC[i] = n + 1 - perm[i]
	}
	return permC
}

func reverse(perm []uint8) []uint8 {
	n := uint8(len(perm))
	permR := make([]uint8, n)
	for i := uint8(0); i < n; i++ {
		permR[n-1-i] = perm[i]
	}
	return permR
}

func inverse(perm []uint8) []uint8 {
	n := uint8(len(perm))
	inv := make([]uint8, n)

	for i := uint8(0); i < n; i++ {
		inv[perm[i]-1] = i + 1
	}

	return inv
}

func GetEqv(perm []uint8) [8][]uint8 {
	var all [8][]uint8

	all[0] = perm
	all[1] = complement(perm) // tl-br diagonal reflect
	all[2] = reverse(perm)    // bl-tr diagonal reflect
	all[3] = reverse(all[1])  // 180 : point reflect

	all[4] = inverse(perm)      // horizontal reflect
	all[5] = complement(all[4]) // left 90
	all[6] = reverse(all[4])    // right 90
	all[7] = reverse(all[5])    // vertical reflect

	return all
}

func IsEqv(perm1 []uint8, perm2 []uint8) bool {
	if len(perm1) != len(perm2) {
		return false
	}
	for _, perm := range GetEqv(perm1) {
		if bytes.Equal(perm, perm2) {
			return true
		}
	}
	return false
}
