package main

import (
	"fmt"
	"strings"
	"sync"
)

type Perm []int

func NewPerm(islice []int) Perm {
	p := make(Perm, 0)
	for _, i := range islice {
		p = append(p, i)
	}
	return p
}

func (p Perm) Add(i int) Perm {
	p = append(p, i)
	return p
}

func (p Perm) Size() int {
	return len(p)
}

func (p Perm) String() string {
	var b strings.Builder
	for _, v := range p {
		fmt.Fprintf(&b, "%d,", v)
	}
	return b.String()
}

// Try with this first, if shit hits the fan
// Try with sync.Map.
type Set struct {
	rwm sync.RWMutex
	set map[string]Perm
}

func NewSet() *Set {
	return &Set{
		set: make(map[string]Perm),
	}
}

func (s *Set) Add(p Perm) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	str := p.String()
	_, prs := s.set[str]
	s.set[str] = p
	return !prs
}

func (s *Set) Get(p Perm) bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	_, prs := s.set[p.String()]
	return prs
}

func (s *Set) Remove(p Perm) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	delete(s.set, p.String())
}

func (s *Set) Size() int {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return len(s.set)
}

func (s *Set) Values() []Perm {
	vals := make([]Perm, 0, s.Size())
	s.rwm.RLock() // Dont push it above the s.Size(), as it is deadlock
	defer s.rwm.RUnlock()

	for _, v := range s.set {
		vals = append(vals, v)
	}
	return vals
}

func localExp(perm Perm, a int) Perm {
	newPerm := make(Perm, 0)

	for _, k := range perm {
		if k < a {
			newPerm = append(newPerm, k)
		} else {
			newPerm = append(newPerm, k+1)
		}
	}
	newPerm = append(newPerm, a)
	return newPerm
}

// go:inline
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isPlane(perm Perm) bool {
	steps := make(Perm, 0)

	for k := 0; k < len(perm)-1; k++ {
		if perm[k] < perm[k+1]-1 {
			steps = append(steps, k)
		}
	}

	for _, s := range steps {
		m, M := perm[s], perm[s+1]
		two, three := 1000, 0
		_ = three // appease the compiler
		prefix, suffix := perm[:s], perm[s+2:]

		for _, k := range prefix {
			if k > m && k < M-1 {
				two = min(k, two)
			}
		}

		for _, k := range suffix {
			if k > two && k < M {
				three = k // I dont know why did the python code init three ????? its returning anyway TODO: check this
				return false
			}
		}
	}
	return true
}

func initCurLevel(s *Set) {
	s.Add(NewPerm([]int{1, 2, 3}))
	s.Add(NewPerm([]int{1, 3, 2}))
	s.Add(NewPerm([]int{2, 1, 3}))
	s.Add(NewPerm([]int{3, 1, 2}))
	s.Add(NewPerm([]int{2, 3, 1}))
	s.Add(NewPerm([]int{3, 2, 1}))
}

var curLevel *Set

func main() {

	curLevel = NewSet()
	initCurLevel(curLevel)
	level := 3

	for level < 20 {
		newLevel := NewSet()
		for _, perm := range curLevel.Values() {
			for a := 1; a < level+2; a++ {
				newPerm := localExp(perm, a)
				if isPlane(newPerm) {
					newLevel.Add(newPerm)
				}
			}

		}
		fmt.Println(newLevel.Size())
		curLevel = newLevel
		level++
	}
}
