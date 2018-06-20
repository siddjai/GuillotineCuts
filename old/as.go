package main

import (
	"fmt"
	"sync"
)

// struct to mimick set in Python
// code from https://play.golang.org/p/_FvECoFvhq
type ArraySet struct {
	set map[[20]int]bool
	mux sync.Mutex
}

func NewArraySet() *ArraySet {
	return &ArraySet{set: make(map[[20]int]bool)}
}

func (set *ArraySet) Add(p [20]int) bool {
	_, found := set.set[p]
	set.set[p] = true
	return !found //False if it existed already
}

func (set *ArraySet) Get(i [20]int) bool {
	_, found := set.set[i]
	return found //true if it existed already
}

func (set *ArraySet) Remove(i [20]int) {
	delete(set.set, i)
}

func main() {
	a := NewArraySet()
	a.Add([20]int{1, 2, 3})
	a.Add([20]int{1, 3, 2})
	a.Add([20]int{2, 1, 3})
	a.Add([20]int{3, 1, 2})
	a.Add([20]int{2, 3, 1})
	a.Add([20]int{3, 2, 1})

	fmt.Println(a.Get([20]int{3, 3, 1}))

	for k := range a.set {
		fmt.Println(k, &k)

	}

}
