package pkg

func mosiacOCS(perm []uint8) uint8 {
	// Separable mosiac
	if IsSeparable(perm) {
		return 0
	}

	// Divisible mosiac
	if cut := IsDivisible(perm); cut > 0 {
		left := perm[:cut]
		for i, p := range left {
			if p > cut {
				left[i] -= cut
			}
		}
		right := perm[cut:]
		for i, p := range right {
			if p > cut {
				right[i] -= cut
			}
		}
		return mosiacOCS(left) + mosiacOCS(right)
	}

	// Non-divisible mosiac
	return 0
}

// Export
func ComputeMosiacOCS(perm []uint8) uint8 {
	return mosiacOCS(perm)
}
