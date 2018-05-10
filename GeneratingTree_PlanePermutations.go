// Plane permutations are those which avoid the pattern 213'54
// Avoiding 213'54 is equivalent to avoiding 2-14-3
// This generating tree based enumeration is based on the following paper:
// https://arxiv.org/pdf/1702.04529.pdf

package main

import "fmt"

// struct to mimick set in Python
// code from https://play.golang.org/p/_FvECoFvhq
type SliceSet struct {
	set map[[]int]bool
}

func NewSliceSet() *SliceSet {
	return &SliceSet{make(map[[]int]bool)}
}

func (set *SliceSet) Add(p []int) bool {
	_, found := set.set[p]
	set.set[p] = true
	return !found	//False if it existed already
}

// func to mimick min in Python
func min (a, b int) (int) {
	if a<=b {return a}
	return b
}

func localExp (perm []int, a int, c chan []int) {
	// Local expansion as described in the paper
	newPerm := make([]int, len(perm)+1)
	for i, k := range perm {
		if k < a {
			newPerm[i] = k
		} else {
			newPerm[i] = k+1
		}
	}
	newPerm[len(perm)] = a
	c <- newPerm
}

func isPlane (perm []int) (bool) {
	n := len(perm)
	steps := make([]int, 0, n)
	for k:=0; k<n-1; k++ {
		if perm[k] < perm[k+1] - 1 {steps = append(steps, k)}
	}

	for _,s := range steps {
		m, M := perm[s], perm[s+1]
		two, three := 1000, 0
		prefix, suffix := perm[:s], perm[s+2:]
		for _,k := range prefix {
			if (k > m) && (k < M - 1) {
				two = min(k, two)
			}
		}

		for _,k := range suffix {
			if (k > two) && (k < M) {
				three = k
				return false
			}
		}
	}

	return true
}

func main() {
	// Implement set - done?
	curLevel := NewSliceSet()
	curLevel.Add([]int{1,2,3})
	curLevel.Add([]int{1,3,2})
	curLevel.Add([]int{2,1,3})
	curLevel.Add([]int{3,1,2})
	curLevel.Add([]int{2,3,1})
	curLevel.Add([]int{3,2,1})
	level := 3
	for level < 20 {
		newLevel := NewSliceSet()
		for _,perm := range curLevel {
			ch := make(chan int)
			for a:=1; a<=level+1; a++ {
				go localExp(perm, a, ch)
			}
			for newPerm := range ch{
				// Implement add func - done?
				if isPlane(newPerm) {newLevel.Add(newPerm)}
			}
			
		}

		fmt.Println(len(newLevel))
		curLevel = newLevel
		level += 1
	}
}
