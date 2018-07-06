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

// O(n^2)
func reflect(perm []uint8) []uint8 {
	rects := BP2FP(perm)
	n := len(rects)
	temp := rects
	remaining := n
	order := make([]uint8, 0)

	for remaining>0 {
		m := uint8(len(temp))
		var blr uint8
		for i:=uint8(0); i<m; i++ {
			r := temp[i]
			if r[0]==0 && r[2]==0 {
				blr = i
				break
			}
		}

		var x [2]uint8
		copy(x[:], temp[blr][:2])
		x_neighbs := make(map[uint8]bool)
		fromTop := true
		for i:=uint8(0); i<m; i++ {
			r := temp[i]
			var intv [2]uint8
			copy(intv[:], r[:2]) 
			if intervalIntersect(intv, x) {
				if r[1] > x[1] {
					fromTop = false
					break
				} else {
					x_neighbs[i] = true
				}
			}
		}

		if !fromTop {
			var y [2]uint8
			copy(y[:], temp[blr][2:])
			y_neighbs := make(map[uint8]bool)
			for i:=uint8(0); i<m; i++ {
				r := temp[i]
				var intv [2]uint8
				copy(intv[:], r[2:])
				if intervalIntersect(intv, y) {
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

		temp[blr] = [4]uint8{255, 255, 255, 255}
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
