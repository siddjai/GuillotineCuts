// This enumeration is based on LTR and RTL succession rule of baxter permutation
// on page 15 of this paper https://arxiv.org/pdf/1702.04529.pdf

package bxt

import (
	"GuillotineCuts/pkg"
)

func EnumNext(perm []uint8) [][]uint8 {
	n := uint8(len(perm))
	perms := make([][]uint8, 0)
	max := uint8(0)
	for i, p := range perm {
		if p > max {
			max = p
			newperm := pkg.InsertPerm(perm, n+1, uint8(i))
			perms = append(perms, newperm)
		}
	}
	max = uint8(0)
	for r := uint8(0); r < n; r++ {
		i := n - r - 1
		p := perm[i]
		if p > max {
			max = p
			newperm := pkg.InsertPerm(perm, n+1, uint8(i+1))
			perms = append(perms, newperm)
		}
	}
	return perms
}
