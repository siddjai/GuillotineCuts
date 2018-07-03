package pkg

// Export
// Baxter permutation to floorplan

func BP2FP(perm []uint8) [][4]uint8 {
	n := uint8(len(perm))
	rects := make([][4]uint8, n+1)
	rects[perm[0]] = [4]uint8{0, n, 0, n}
	below := make(map[uint8]uint8)
	left := make(map[uint8]uint8)
	prevlabel := perm[0]

	for k := uint8(1); k < n; k++ {
		p := perm[k]
		if p < prevlabel {
			oldrect := rects[prevlabel]
			// middle := (oldrect[2] + oldrect[3]) / 2

			// Horizontal slice
			rects[p] = oldrect
			rects[p][2] = k
			rects[prevlabel][3] = k

			// Store spatial relations
			below[p] = prevlabel
			lp, past := left[prevlabel]
			if past {
				left[p] = lp
			}

			_, ok := left[p]
			for ok && left[p] > p {
				l := left[p]

				rects[p][0] = rects[l][0]

				rects[l][3] = rects[p][2]

				ll, okl := left[l]
				if okl {
					left[p] = ll
				} else {
					delete(left, p)
				}

				ok = okl
			}

			prevlabel = p

		} else {
			oldrect := rects[prevlabel]
			// middle := (oldrect[0] + oldrect[1]) / 2

			// Vertical slice
			rects[p] = oldrect
			rects[p][0] = k
			rects[prevlabel][1] = k

			// Store spatial relations
			left[p] = prevlabel
			bp, past := below[prevlabel]
			if past {
				below[p] = bp
			}

			_, ok := below[p]
			for ok && below[p] < p {
				b := below[p]

				rects[p][2] = rects[b][2]

				rects[b][1] = rects[p][0]

				bb, okb := below[b]
				if okb {
					below[p] = bb
				} else {
					delete(below, p)
				}

				ok = okb
			}

			prevlabel = p

		}
	}

	return rects
}
