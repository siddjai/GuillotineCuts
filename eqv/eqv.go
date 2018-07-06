package eqv

import(
	"GuillotineCuts/BP2FP"
)

func intervalIntersect(i1 [2]uint8, i2 [2]uint8) bool{
	return !(i1[0]>=i2[1] || i2[0]>=i1[1])
}

func complement(perm []uint8) []uint8 {
	n := uint8(len(perm))
	permC := make([]uint8, n)
	for i:=uint8(0); i<n; i++ {
		permC[i] = n + 1 - perm[i]
	}
	return permC
}

func reverse(perm []uint8) []uint8 {
	n := uint8(len(perm))
	permR := make([]uint8, n)
	for i:=uint8(0); i<n; i++ {
		permR[n-1-i] = perm[i]
	}
	return permR
}

func inversePerm(perm []uint8) []uint8 {
	n := uint8(len(perm))
	inv := make([]uint8, n)

	for i:=uint8(0); i<n; i++ {
		inv[perm[i]-1] = i+1
	}

	return inv
}


func eqv(perm []uint8) [8][]uint8 {
	var all [8][]uint8

	all[0] = perm
	all[1] = complement(perm)
	all[2] = reverse(perm)
	all[3] = reverse(all[1])

	all[4] = inversePerm(perm)
	all[5] = complement(all[4])
	all[6] = reverse(all[4])
	all[7] = reverse(all[5])

	return all
}
