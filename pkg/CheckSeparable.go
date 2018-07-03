package pkg

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

func addRange(stack *[][2]uint8, r [2]uint8) {
	s := *stack
	n := len(s)
	if n == 0 {
		*stack = append(s, r)
		return
	}

	top := s[n-1]
	if r[0] > top[1]+1 || top[0] > r[1]+1 {
		*stack = append(s, r)
		return
	} else {
		*stack = s[:n-1]
		r_new := [2]uint8{Min(r[0], top[0]), Max(r[1], top[1])}
		addRange(stack, r_new)
	}
}

// Export
func IsSeparable(perm []uint8) bool {
	if len(perm) <= 4 {
		return true
	}

	var stack [][2]uint8
	for _, p := range perm {
		r := [2]uint8{p, p}
		addRange(&stack, r)
	}
	return len(stack) == 1
}
