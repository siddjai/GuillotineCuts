package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"strings"
	"time"
)

const maxPerm = 20

var (
	maxLevel = flag.Int("l", 10, "MAX Level")
	p        = flag.Bool("pprof", false, "Enable Profiling")
	t        = flag.Bool("trace", false, "Enable Tracing")
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
	b.Grow(32)
	for _, v := range p {
		b.WriteString(string(v))
	}
	return b.String()
}

// Try with this first, if shit hits the fan
// Try with sync.Map.
type Set struct {
	set map[string]Perm
}

func NewSet() *Set {
	return &Set{
		set: make(map[string]Perm),
	}
}

func (s *Set) Add(p Perm) bool {
	str := p.String()
	_, prs := s.set[str]
	s.set[str] = p
	return !prs
}

func (s *Set) Get(p Perm) bool {
	_, prs := s.set[p.String()]
	return prs
}

func (s *Set) Remove(p Perm) {
	delete(s.set, p.String())
}

func (s *Set) Size() int {
	return len(s.set)
}

func (s *Set) Values() []Perm {
	vals := make([]Perm, 0, s.Size())

	for _, v := range s.set {
		vals = append(vals, v)
	}
	return vals
}

func localExp(perm Perm, a int) Perm {
	newPerm := make(Perm, len(perm)+1)

	var i int
	for i = 0; i < len(perm); i++ {
		k := perm[i]
		if k < a {
			newPerm[i] = k
		} else {
			newPerm[i] = k + 1
		}
	}
	newPerm[i] = a
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
	steps := make(Perm, len(perm)-1)

	var i int
	for k := 0; k < len(perm)-1; k++ {
		if perm[k] < perm[k+1]-1 {
			steps[i] = k
			i++
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
	flag.Parse()
	defer trackTime(time.Now(), "MAIN")

	if *p {
		log.Println("Profiling Enabled")
		pf, err := os.Create("pprof.out")
		if err != nil {
			log.Fatal("Could not create pprof file")
		}
		defer pf.Close()
		pprof.StartCPUProfile(pf)
		defer pprof.StopCPUProfile()
	}

	if *t {
		log.Println("Tracing Enabled")
		tf, err := os.Create("trace.out")
		if err != nil {
			log.Fatal("Could not create trace file")
		}
		defer tf.Close()
		trace.Start(tf)
		defer trace.Stop()
	}

	curLevel = NewSet()
	initCurLevel(curLevel)
	level := 3

	for level < *maxLevel {
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

func trackTime(s time.Time, msg string) {
	fmt.Println(msg, ":", time.Since(s))
}
