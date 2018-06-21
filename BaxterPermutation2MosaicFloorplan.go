// Given a Baxter permutation, this program constructs a corresponding floorplan
// Based on the mapping mentioned on page 15 in this thesis:
// https://www.cs.technion.ac.il/users/wwwb/cgi-bin/tr-get.cgi/2006/PHD/PHD-2006-11.pdf
// And the related paper
// Eyal Ackerman, Gill Barequet, and Ron Y. Pinter.  A bijection
// between permutations and floorplans, and its applications.
// Discrete Applied Mathematics, 154(12):1674–1684, 2006.

// Format for specifying a rectangle:
// (x1, x2, y1, y2) where x1 is the min and x2 is the max y coordinate.
// Similarly for y

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	perm := make([]int, n)
	for i:=0; i<n; i++ {
		var k int
		fmt.Scanf("%d ", &k)
		perm[i] = k
	}

	rects := make(map[int][4]int)
	rects[perm[0]] = [4]int{0,n,0,n}
	below = make(map[int]int)
	left = make(map[int]int)
	prevlabel := perm[0]

	for k:=1; k<n; k++ {
		p := perm[k]
		if p < prevlabel {
			oldrect := rects[prevlabel]
			middle := (oldrect[2] + oldrect[3])/2
			rects[p] = oldrect
			rects[p][2] = middle
			rects[prevlabel][3] = middle
			below[p] = prevlabel

			//Implement for (while)

			prevlabel = p

		} else {
			oldrect := rects[prevlabel]
			middle := (oldrect[0]+oldrect[1])/2
			rects[p] = oldrect
			rects[p][0] = middle
			rects[prevlabel][1] = middle
			left[p] = prevlabel

			//Implement for (while)

			prevlabel = p

		}
	}

	fin_rects := make([][4]int, len(rects))
    idx := 0
    for  _, rec := range rects {
       fin_rects[idx] = rec
       idx++
    }

    fmt.Println(fin_rects)
}