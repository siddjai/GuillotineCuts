package BP2FP

import (
	"fmt"
	"strconv"
	"strings"
)

func draw(rects [][4]uint8) {
	n := len(rects)
	mat := make([][]uint8, n)
	for k := 0; k < n; k++ {
		mat[k] = make([]uint8, n)
	}

	for k, rec := range rects {
		for i := rec[0]; i < rec[1]; i++ {
			for j := rec[2]; j < rec[3]; j++ {
				if mat[n-1-int(j)][i] != 0 {
					fmt.Println("Overwrite!")
				}
				mat[n-1-int(j)][i] = uint8(k + 1)
			}
		}
	}

	for k := 0; k < n; k++ {
		for l := 0; l < n; l++ {
			fmt.Printf("%d", mat[k][l])
		}
		fmt.Println()
	}
	fmt.Println()
}

func OnlineTest() {
	var n int
	fmt.Scanf("%d\n", &n)
	perm := make([]uint8, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&perm[i])
	}

	rects := BP2FP(perm)

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

func OfflineTest(permStr string) {
	tokens := strings.Split(permStr, " ")
	n := len(tokens)
	perm := make([]uint8, len(tokens))
	for i := 0; i < len(tokens); i++ {
		i64, _ := strconv.ParseUint(tokens[i], 10, 8)
		perm[i] = uint8(i64)
	}
	rects := BP2FP(perm)

	draw(rects)

	fmt.Println(n)
	for k := 0; k < n; k++ {
		for i, rec := range rects[k] {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Print(rec)
		}
		fmt.Println()
	}
}
