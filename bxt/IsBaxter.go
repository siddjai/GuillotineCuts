package bxt

import "GuillotineCuts/pkg"

// O(n^2)
func IsBaxter(perm []uint8) bool {

	for k := 0; k < len(perm)-1; k++ {

		two, three := uint8(len(perm)+1), uint8(0)

		if perm[k] < perm[k+1]-2 {

			// Memorise -14-
			m, M := perm[k], perm[k+1]
			prefix, suffix := perm[:k], perm[k+2:]

			//Avoid 3-14-2
			for _, k := range prefix {
				if (k > m+1) && (k < M) {
					three = pkg.Max(k, three)
				}
			}

			for _, k := range suffix {
				if (k < three) && (k > m) {
					two = k
					return false
				}
			}
		}

		if perm[k] > perm[k+1]+2 {

			// Memorise -41-
			m, M := perm[k+1], perm[k]
			prefix, suffix := perm[:k], perm[k+2:]

			// Avoid 2-14-3
			for _, k := range prefix {
				if k > m && k < M-1 {
					two = pkg.Min(k, two)
				}
			}

			for _, k := range suffix {
				if k > two && k < M {
					three = k
					return false
				}
			}
		}
	}

	return true
}
