package pkg

func toPermutation(arr []uint8) []uint8 {
	min := uint8(len(arr) + 1)
	for _, p := range arr {
		min = Min(min, p)
	}
	res := arr[:]
	if min > 1 {
		for i := 0; i < len(arr); i++ {
			res[i] -= min - 1
		}
	}
	return res
}

func mosiacOCS(perm []uint8) uint8 {
	// Separable mosiac
	if IsSeparable(perm) {
		return 0
	}

	// Divisible mosiac
	if cut := IsMosaicDivisible(perm); cut > 0 {
		left := toPermutation(perm[:cut])
		right := toPermutation(perm[cut:])
		return mosiacOCS(left) + mosiacOCS(right)
	}

	// Non-divisible mosiac

	return 0
}

// Export
func ComputeMosiacOCS(perm []uint8) uint8 {
	return mosiacOCS(perm)
}
