// Baxter permutations are those which avoid the patterns 2-41-3 and 3-14-2 
// This generating tree based enumeration is based on the following paper:
// https://arxiv.org/pdf/1702.04529.pdf

// Format for specifying a rectangle:
// (x1, x2, y1, y2) where x1 is the min and x2 is the max y coordinate.
// Similarly for y

// Format for specifying a line:
// (a, b) where a is a number and b is a binary number where
// 0 -> || to Y axis AND 1 -> || to X axis

// Format for specifying an interval:
// (a1, a2)

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
	"sort"
)

var (
	maxLevel = flag.Int("l", 6, "MAX Level")
	procs    = flag.Int("procs", 2, "Number of workers")
	p        = flag.Bool("pprof", false, "Enable Profiling")
	t        = flag.Bool("trace", false, "Enable Tracing")
)

// Compute Optimal Cut Seq

// var dp_seq map[[4]int][][6]int
// var dp_kill map[[4]int]int

func intervalIntersect(i1 [2]int, i2 [2]int) (bool){
	if i1[0]>=i2[1] || i2[0]>=i1[1] {
		return false
	}
	return true
}


func optimalCut(rects [][4]int, x []int, y []int, reg [4]int, seq [][6]int, dp_kill map[[4]int]int, dp_seq map[[4]int][][6]int) ([][6]int, int){
	// rects : Rectangles in the current set
	// x : sorted list of X coordinates
	// y : sorted list of Y coordinates
	// code : labels of current set encoded
	// seq : sequence of cuts upto this set; (code, value, orientation)

	// RETURN
	// seq : seq of rectangles including this level
	// killed : no of rectangles killed including this level

	// Choosing to not compute cuts for small enough sets
	if len(rects) <= 3 {
		return seq, 0
	}

	// Check if stored
	killed, ok := dp_kill[reg]
	if ok {
		sseq := dp_seq[reg]
		return sseq, killed
	}

	m := len(x) + len(y) - 4
	cuts := make([]int, m)
	seqs := make(map[int][][6]int)

	for i:=0; i<m; i++ {
		//A high enough constant
		cuts[i] = 255
	}

	//Enough to try all rectangle edges
	for k:=0; k<len(x)-2; k++ {
		var rects1 [][4]int
		var rects2 [][4]int
		boundary := false

		kill_cur := 0
		for _, rec := range rects {
			var xi [2]int
			xi[0] = rec[0]
			xi[1] = rec[1]
			if intervalIntersect(xi, [2]int{x[1+k], x[1+k]}) {
				kill_cur++
			} else if rec[1] <= x[1+k] {
				rects1 = append(rects1, rec)
			} else {
				rects2 = append(rects2, rec)
			}

			if rec[0] == x[1+k] {boundary = true}

		}

		xx1 := x[:2+k]
		xx2 := x[2+k:]
		if boundary { xx2 = append([]int{x[1+k]}, xx2...) }

		yy1m := make(map[int]bool)
		for _, rec := range rects1 {
			yy1m[rec[2]] = true
			yy1m[rec[3]] = true
		}

		var reg1 [4]int
		reg1 = reg 
		reg1[1] = x[1+k]

		yy2m := make(map[int]bool)
		for _, rec := range rects2 {
			yy2m[rec[2]] = true
			yy2m[rec[3]] = true
		}

		var reg2 [4]int
		reg2 = reg
		reg2[0] = xx2[0]

		yy1 := make([]int, len(yy1m))
		yy2 := make([]int, len(yy2m))

		i := 0
		for k := range yy1m {
		    yy1[i] = k
		    i++
		}

		i = 0
		for k := range yy2m {
		    yy2[i] = k
		    i++
		}

		sort.Ints(yy1)
		sort.Ints(yy2)

		//Jugaad | Should no longer be needed
		kill1 := 255
		kill2 := 255
		var seq1 [][6]int
		var seq2 [][6]int

		if len(rects1) < len(rects) && len(rects2) < len(rects) {
			seq1, kill1 = optimalCut(rects1, xx1, yy1, reg1, seq, dp_kill, dp_seq)
			seq2, kill2 = optimalCut(rects2, xx2, yy2, reg2, seq, dp_kill, dp_seq)
		}

		cuts[k] = kill1 + kill2 + kill_cur

		seq = append(seq, seq1...)
		seq = append(seq, seq2...)
		seqs[k] = seq

	}

	for k:=0; k<len(y)-2; k++ {
		var rects1 [][4]int
		var rects2 [][4]int
		boundary := false
		kill_cur := 0
		for _, rec := range rects {
			var yi [2]int 
			yi[0] = rec[2]
			yi[1] = rec[3]
			if intervalIntersect(yi, [2]int{y[1+k], y[1+k]}) {
				kill_cur++
			} else if rec[3] <= y[1+k] {
				rects1 = append(rects1, rec)
			} else {
				rects2 = append(rects2, rec)
			}

			if rec[2] == y[1+k] {boundary = true}
		}

		yy1 := y[:2+k]
		yy2 := y[2+k:]
		if boundary { yy2 = append([]int{y[1+k]}, yy2...) }

		xx1m := make(map[int]bool)
		for _, rec := range rects1 {
			xx1m[rec[0]] = true
			xx1m[rec[1]] = true 
		}

		var reg1 [4]int
		reg1 = reg 
		reg1[3] = y[1+k]

		xx2m := make(map[int]bool) 
		for _, rec := range rects2 {
			xx2m[rec[0]] = true
			xx2m[rec[1]] = true
		}

		var reg2 [4]int
		reg2 = reg
		reg2[2] = yy2[0]

		xx1 := make([]int, len(xx1m))
		xx2 := make([]int, len(xx2m))

		i := 0
		for k := range xx1m {
		    xx1[i] = k
		    i++
		}

		i = 0
		for k := range xx2m {
		    xx2[i] = k
		    i++
		}

		sort.Ints(xx1)
		sort.Ints(xx2)

		//Jugaad | should no longer be needed
		kill1 := 255
		kill2 := 255
		var seq1 [][6]int
		var seq2 [][6]int

		if len(rects1) < len(rects) && len(rects2) < len(rects) {
			seq1, kill1 = optimalCut(rects1, xx1, yy1, reg1, seq, dp_kill, dp_seq)
			seq2, kill2 = optimalCut(rects2, xx2, yy2, reg2, seq, dp_kill, dp_seq)
		}

		cuts[len(x) - 2 + k] = kill1 + kill2 + kill_cur

		seq = append(seq, seq1...)
		seq = append(seq, seq2...)
		seqs[len(x) - 2 + k] = seq
	}

	//Choose min
	minPtr := 0
	for k:=0; k<m; k++ {
		if cuts[k] < cuts[minPtr] {minPtr = k}
	}

	newLine := [2]int{1000, 0}
	if minPtr < len(x) - 2 {
		newLine = [2]int{x[1+minPtr], 0}
	} else {
		newLine = [2]int{y[minPtr - len(x) + 2], 1}
	}

	dp_kill[reg] = cuts[minPtr]
	var cur [][6]int
	cur = append(cur, [6]int{reg[0], reg[1], reg[2], reg[3], newLine[0], newLine[1]})
	best_seq := seqs[minPtr]
	seqf := append(cur, best_seq...)
	dp_seq[reg] = seqf

	return seqf, cuts[minPtr]

}

func sanityCheck(rects [][4]int) (bool){
	for _, rec1 := range rects {
		for _, rec2 := range rects {
			x1 := [2]int{rec1[0], rec1[1]}
			x2 := [2]int{rec2[0], rec2[1]}
			y1 := [2]int{rec1[2], rec1[3]}
			y2 := [2]int{rec2[2], rec2[3]}
			if intervalIntersect(x1, x2) && intervalIntersect(y1, y2) {
				return false
			}
		}
	}

	return true
}

func ComputeOCS(rects [][4]int) ([][6]int, int) {
	dp_seq := make(map[[4]int][][6]int)
	dp_kill := make(map[[4]int]int)

	if sanityCheck(rects) {
		xm := make(map[int]bool)
		ym := make(map[int]bool)
		for _, tup := range rects {
			xm[tup[0]] = true 
			xm[tup[1]] = true
			ym[tup[2]] = true 
			ym[tup[3]] = true 
		}

		x := make([]int, len(xm))
		y := make([]int, len(ym))

		i := 0
		for k := range xm {
		    x[i] = k
		    i++
		}

		i = 0
		for k := range ym {
		    y[i] = k
		    i++
		}		

		sort.Ints(x)
		sort.Ints(y)
		reg := [4]int{x[0], x[len(x)-1], y[0], y[len(y)-1]}
		var seq [][6]int

		return optimalCut(rects, x, y, reg, seq, dp_kill, dp_seq)

	} else {
		fmt.Println("Invalid set!")
		return nil, 0
	}
}

// End of ComputeOCS



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

// Given a Baxter permutation, the function BP2FP constructs a corresponding floorplan
// Based on the mapping mentioned on page 15 in this thesis:
// https://www.cs.technion.ac.il/users/wwwb/cgi-bin/tr-get.cgi/2006/PHD/PHD-2006-11.pdf
// And the related paper
// Eyal Ackerman, Gill Barequet, and Ron Y. Pinter.  A bijection
// between permutations and floorplans, and its applications.
// Discrete Applied Mathematics, 154(12):1674–1684, 2006.

func BP2FP(perm Perm, n int) [][4]int {
	
	rects := make([][4]int, n+1)
	rects[perm[0]] = [4]int{0, n, 0, n}
	below := make(map[int]int)
	left := make(map[int]int)
	prevlabel := perm[0]

	for k := 1; k < n; k++ {
		p := perm[k]
		if p < prevlabel {
			oldrect := rects[prevlabel]
			// middle := (oldrect[2] + oldrect[3]) / 2

			// Horizontal slice
			rects[p] = oldrect
			rects[p][2] = k
			rects[prevlabel][3] = k

			// Store spatial relations
			below[p] = prevlabel
			lp, past := left[prevlabel]
			if past {
				left[p] = lp
			}

			_, ok := left[p]
			for ok && left[p] > p {
				l := left[p]

				rects[p][0] = rects[l][0]

				rects[l][3] = rects[p][2]

				ll, okl := left[l]
				if okl {
					left[p] = ll
				} else {
					delete(left, p)
				}

				ok = okl
			}

			prevlabel = p

		} else {
			oldrect := rects[prevlabel]
			// middle := (oldrect[0] + oldrect[1]) / 2

			// Vertical slice
			rects[p] = oldrect
			rects[p][0] = k
			rects[prevlabel][1] = k

			// Store spatial relations
			left[p] = prevlabel
			bp, past := below[prevlabel]
			if past {
				below[p] = bp
			}

			_, ok := below[p]
			for ok && below[p] < p {
				b := below[p]

				rects[p][2] = rects[b][2]

				rects[b][1] = rects[p][0]

				bb, okb := below[b]
				if okb {
					below[p] = bb
				} else {
					delete(below, b)
				}

				ok = okb
			}

			prevlabel = p

		}
		//draw(rects, n)
	}

	return rects[1:]
}
// End of BP2FP

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

func addRange(stack *[][2]int, r [2]int) (){
	s := *stack
	n := len(s)
	if n==0 {
		*stack = append(s, r)
		return
	}
	
	top := s[n-1]
	if r[0]>top[1]+1 || top[0]>r[1]+1 {
		*stack = append(s, r)
		return
	} else {
		*stack = s[:n-1]
		r_new := [2]int{min(r[0], top[0]), max(r[1], top[1])}
		addRange(stack, r_new)
	}
}

func isSeperable(perm Perm) bool {
    var stack [][2]int
    for _, p := range perm {
    	r := [2]int{p, p}
    	addRange(&stack, r)
    }

    if len(stack)==1 {
    	return true
    }
    return false

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
			if !isSeperable(newPerm) {
				n := level+1
				fmt.Println(n)
				//fmt.Println(newPerm)
				rects := BP2FP(newPerm, n)
				_, kill := ComputeOCS(rects)

				lock.Lock()
				levelPermCount[n]++
				if kill >= n/4 {
					// Save to file instead
					fmt.Println(newPerm)
					fmt.Println()
				}
				lock.Unlock()
			}
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