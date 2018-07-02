// This algorithm is taken from the following paper
// https://www.sciencedirect.com/science/article/pii/S0020019097002093

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func addRange(stack *[][2]int, r [2]int) {
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
		r_new := [2]int{min(r[0], top[0]), max(r[1], top[1])}
		addRange(stack, r_new)
	}
}

func IsSeparable(perm []int) bool {
	var stack [][2]int
	for _, p := range perm {
		r := [2]int{p, p}
		addRange(&stack, r)
	}
	return len(stack) == 1
}

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
	perm := make([]int, len(tokens))
	for i := 0; i < len(tokens); i++ {
		perm[i], _ = strconv.Atoi(tokens[i])
	}

	fmt.Println(IsSeparable(perm))
}
