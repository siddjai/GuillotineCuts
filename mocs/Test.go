package mocs

import (
	"GuillotineCuts/bxt"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func IndividualTest(permStr string) {
	tokens := strings.Split(permStr, " ")
	perms := make([][]uint8, 0)
	perm := make([]uint8, len(tokens))
	for i := 0; i < len(tokens); i++ {
		i64, _ := strconv.ParseUint(tokens[i], 10, 8)
		perm[i] = uint8(i64)
	}
	perms = append(perms, perm)
	n := len(perm)
	m := NewManager(uint8(n), 1, 1)

	go m.AddJobs(perms)
	go m.StopJobs()
	m.Start()
	m.PrintResult(false)
}

func BatchTest(n int, N uint32) {
	m := NewManager(uint8(n), 100, 1000)
	perms := make([][]uint8, 0)
	for i := uint32(0); i < N; i++ {
		perm := make([]uint8, n)
		for j, p := range rand.Perm(n) {
			perm[j] = uint8(p + 1)
		}
		if bxt.IsBaxter(perm) {
			perms = append(perms, perm)
		}
	}
	fmt.Println("Total perms:", len(perms))

	go m.AddJobs(perms)
	go m.StopJobs()
	m.Start()
	m.PrintResult(false)
}
