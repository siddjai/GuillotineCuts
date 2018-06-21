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
	maxLevel = flag.Int("l", 9, "MAX Level")
	procs    = flag.Int("procs", 2, "Number of workers")
	p        = flag.Bool("pprof", false, "Enable Profiling")
	t        = flag.Bool("trace", false, "Enable Tracing")
)

type Perm []int

type Data struct {
	perm  Perm
	level int
}

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
var lock sync.RWMutex

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

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	jobs := make(chan *Data, 100)
	results := make(chan *Data, 100)

	for w := 1; w <= *procs; w++ {
		go worker(jobs, results)
	}

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for _, perm := range curLevel.Values() {
		jobs <- &Data{perm: perm, level: 3}
	}


	// Finally we collect all the results of the work and give make jobs
	for data := range results{
		fmt.Println(levelPermCount)
		if data.level == *maxLevel {
			close(jobs)
			break
		}
		jobs <- data
	}

	//fmt.Println(levelPermCount)

}

func worker(jobs <-chan *Data, results chan<- *Data) {

	for perm := range jobs {

		if perm.level >= *maxLevel {
			return
		}
		for a := 1; a < perm.level+2; a++ {
			newPerm := localExp(perm.perm, a)
			if isPlane(newPerm) {
				lock.Lock()
				levelPermCount[perm.level]++
				lock.Unlock()
				newPermData := Data{perm: newPerm, level: perm.level + 1}
				results <- &newPermData
			}
		}
	}
}

func trackTime(s time.Time, msg string) {
	fmt.Println(msg, ":", time.Since(s))
}
