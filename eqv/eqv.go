package eqv

import(
	"GuillotineCuts/BP2FP"
)

func intervalIntersect(i1 [2]int, i2 [2]int) {
	return !(i1[0]>=i2[1] || i2[0]>=i1[1])
}

func complement(perm []uint8) []uint8 {
	n := uint8(len(perm))
	permC := make([]uint8, n)
	for i, k := range perm {
		permC[i] = n + 1 - k
	}
	return permC
}

func reverse(perm []uint8) []uint8 {
	n := uint8(len(perm))
	permR := make([]uint8, n)
	for i, k := range perm {
		permR[n-1-i] = k
	}
	return permR
}

func inversePerm(perm []uint8) []uint8 {
	n := len(perm)
	inv := make([]int, n)

	for i:=0; i<n; i++ {
		inv[perm[i]-1] = i+1
	}

	return inv
}

// O(n^2)
func reflect(perm []uint8) []uint8 {
	rects := BP2FP.BP2FP(perm)
	n := len(rects)
	temp := rects
	remaining := n
	order := make([]uint8, 0)

	for remaining>0 {
		blr := -1
		var rblr [4]uint8
		for i, r := range temp {
			if r[0]==0 && r[2]==0 {
				blr = i
				break
			}
		}
		rblr = temp[blr]

		x := temp[blr][:2]
		x_neighbs := make(map[int]bool)
		fromTop := true
		for i, r := range temp {
			if intervalIntersect(r[:2], x) {
				if r[1] > x[1] {
					fromTop = false
					break
				} else {
					x_neighbs[i] = true
				}
			}
		}

		if !fromTop {
			y := temp[blr][2:]
			y_neighbs := make(map[int]bool)
			for i, r := range temp {
				if intervalIntersect(r[2:], y) {
					y_neighbs[i] = true
				}
			}

			for i := range y_neighbs {
				temp[i][0] = 0
			}

		} else {
			for i := range x_neighbs {
				temp[i][2] = 0
			}

		}

		temp[blr] = [4]int{-1, -1, -1, -1}
		remaining--
		order = append(order, blr+1)
	}

	return inversePerm(order)
	
}

func eqv(perm []uint8) [8][]uint8 {
	var all [8][]uint8

	all[0] = perm
	all[1] = complement(perm)
	all[2] = reverse(perm)
	all[3] = reverse(all[1])

	all[4] = reflect(perm)
	all[5] = complement(all[4])
	all[6] = reverse(all[4])
	all[7] = reverse(all[5])

	return all
}