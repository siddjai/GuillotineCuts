package bxt

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

	fmt.Println(IsBaxter(perm))
}

func OfflineTest(permStr string) {
	tokens := strings.Split(permStr, " ")
	perm := make([]uint8, len(tokens))
	for i := 0; i < len(tokens); i++ {
		i64, _ := strconv.ParseUint(tokens[i], 10, 8)
		perm[i] = uint8(i64)
	}

	fmt.Println(IsBaxter(perm))
}
