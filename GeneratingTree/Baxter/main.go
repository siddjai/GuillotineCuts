// Baxter permutations are those which avoid the patterns 2-41-3 and 3-14-2 
// This generating tree based enumeration is based on the following paper:
// https://arxiv.org/pdf/1702.04529.pdf

package main

import (
	"fmt"
	"strings"
	"sync"
	"flag"
	"time"
	"os"
	"runtime/trace"
	"log"
	"runtime/pprof"
)

var (
	maxLevel = flag.Int("l", 11, "MAX Level")
	procs    = flag.Int("procs", 2, "Number of workers")
	p        = flag.Bool("pprof", false, "Enable Profiling")
	t        = flag.Bool("trace", false, "Enable Tracing")
)

// Level at which the Tree formation starts
const startingLevel  = 4
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

var levelPermCount = make(map[int]int)
var lock sync.Mutex

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
	newPerm := make(Perm, 0, len(perm)+1)

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBaxter(perm Perm) bool {

	for k := 0; k < len(perm)-1; k++ {

		two, three := 1000, 0

		if perm[k] < perm[k+1]-2 {

			// Memorise -14-
			m, M := perm[k], perm[k+1]
			prefix, suffix := perm[:k], perm[k+2:]

			two = 1000
			three = 0

			//Avoid 3-14-2
			for _, k := range prefix {
				if (k > m + 1) && (k < M) {
					three = max(k, three)
				}
			}

			for _, k := range suffix {
				if (k < three) && (k > m){
					two = k
					return false
				}
			}
		}

		if perm[k] > perm[k+1]+2 {

			// Memorise -41-
			m, M:= perm[k+1], perm[k]
			prefix, suffix := perm[:k], perm[k+2:]

			// Avoid 2-14-3
			for _, k := range prefix {
				if k > m && k < M-1 {
					two = min(k, two)
				}
			}

			for _, k := range suffix {
				if k > two && k < M {
					three = k
					return false
				}
			}
		}
	}

	return true
}

func initCurLevel(s *Set) {
	p := NewSet()
	p.Add(NewPerm([]int{1, 2, 3}))
	p.Add(NewPerm([]int{1, 3, 2}))
	p.Add(NewPerm([]int{2, 1, 3}))
	p.Add(NewPerm([]int{3, 1, 2}))
	p.Add(NewPerm([]int{2, 3, 1}))
	p.Add(NewPerm([]int{3, 2, 1}))

	for _, perm := range p.Values() {
		for a:=1; a<=4; a++ {
			newPerm := localExp(perm, a)
			if isBaxter(newPerm) {
				s.Add(newPerm)
			}
		}
	}
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

	var wg sync.WaitGroup
	for _, perm := range curLevel.Values() {
		wg.Add(1)
		go worker(perm, startingLevel, &wg)
	}

	wg.Wait()
	fmt.Println(levelPermCount)

}

func worker(perm Perm, level int, wg *sync.WaitGroup) {

	if level >= *maxLevel {
		return
	}
	for a := 1; a <= level+1; a++ {
		newPerm := localExp(perm, a)
		if isBaxter(newPerm) {
			lock.Lock()
			levelPermCount[level+1]++
			lock.Unlock()
			worker(newPerm, level+1, wg)
		}
	}

	if level == startingLevel {
		wg.Done()
	}

}

func trackTime(s time.Time, msg string) {
	fmt.Println(msg, ":", time.Since(s))
}