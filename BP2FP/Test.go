package BP2FP

import (
	"fmt"
	"strconv"
	"strings"
)

func OnlineTest() {
	var n int
	fmt.Scanf("%d\n", &n)
	perm := make([]uint8, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&perm[i])
	}

	rects := BP2FP(perm)

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
