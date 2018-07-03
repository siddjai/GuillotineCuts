// This algorithm is taken from the following paper
// https://www.sciencedirect.com/science/article/pii/S0020019097002093

package main

import (
	pkg "GuillotineCuts/pkg"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Online
	// var n int
	// fmt.Scanf("%d\n", &n)
	// perm := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&perm[i])
	// }

	// Offline
	permStr := "3 2 1 4 13 12 7 8 11 10 9 6 5"
	tokens := strings.Split(permStr, " ")
	perm := make([]uint8, len(tokens))
	for i := 0; i < len(tokens); i++ {
		i64, _ := strconv.ParseUint(tokens[i], 10, 8)
		perm[i] = uint8(i64)
	}

	fmt.Println(pkg.IsSeparable(perm))
}
