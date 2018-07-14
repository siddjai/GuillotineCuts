package eqv

import (
	"fmt"
	"strconv"
	"strings"
)

func PairTest(permstr1 string, permstr2 string) {
	tokens := strings.Split(permstr1, " ")
	perm1 := make([]uint8, len(tokens))
	for i := 0; i < len(tokens); i++ {
		i64, _ := strconv.ParseUint(tokens[i], 10, 8)
		perm1[i] = uint8(i64)
	}
	tokens = strings.Split(permstr2, " ")
	perm2 := make([]uint8, len(tokens))
	for i := 0; i < len(tokens); i++ {
		i64, _ := strconv.ParseUint(tokens[i], 10, 8)
		perm2[i] = uint8(i64)
	}
	fmt.Println(IsEqv(perm1, perm2))
}
