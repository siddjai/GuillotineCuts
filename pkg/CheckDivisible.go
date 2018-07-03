package pkg

// Export
// Return 0 if not divisible
// Otherwise return the cut from 1 to n-1

func IsDivisible(perm []uint8) uint8 {
	n := uint8(len(perm))
	if perm[0] == 1 || perm[0] == n {
		return 1
	}
	if perm[n-1] == 1 || perm[n-1] == n {
		return n - 1
	}

	var i1, in uint8
	for i, p := range perm {
		if p == 1 {
			i1 = uint8(i)
		}
		if p == n {
			in = uint8(i)
		}
	}

	if i1 < in {
		max := perm[0]
		for i := uint8(1); i < in; i++ {
			max = Max(max, perm[i])
			if max == i+1 {
				return i + 1
			}
		}
	} else {
		min := perm[0]
		for i := uint8(1); i < i1; i++ {
			min = Min(min, perm[i])
			if min+i == n {
				return i + 1
			}
		}
	}
	return 0
}
