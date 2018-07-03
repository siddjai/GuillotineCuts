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
	pkg "GuillotineCuts/pkg"
	"fmt"
	"strconv"
	"strings"
)

func draw(rects [][4]uint8) {
	n := uint8(len(rects))
	rects = rects[1:]
	mat := make([][]uint8, n)
	for k := uint8(0); k < n; k++ {
		mat[k] = make([]uint8, n)
	}

	for k, rec := range rects {
		for i := rec[0]; i < rec[1]; i++ {
			for j := rec[2]; j < rec[3]; j++ {
				if mat[n-1-j][i] != 0 {
					fmt.Println("Overwrite!")
				}
				mat[n-1-j][i] = uint8(k + 1)
			}
		}
	}

	for k := uint8(0); k < n; k++ {
		for l := uint8(0); l < n; l++ {
			fmt.Printf("%d", mat[k][l])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	// Online
	// var n int
	// fmt.Scanf("%d\n", &n)
	// perm := make([]uint8, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&perm[i])
	// }

	//Offline
	// Offline
	permStr := "2 8 6 4 3 5 1 7"
	tokens := strings.Split(permStr, " ")
	n := len(tokens)
	perm := make([]uint8, len(tokens))
	for i := 0; i < len(tokens); i++ {
		i64, _ := strconv.ParseUint(tokens[i], 10, 8)
		perm[i] = uint8(i64)
	}
	rects := pkg.BP2FP(perm)
	draw(rects)

	fmt.Println(n)
	for k := 1; k <= n; k++ {
		for i, rec := range rects[k] {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Print(rec)
		}
		fmt.Println()
	}
}
