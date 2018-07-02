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
	"sort"
)

var dp_seq map[[4]int][][6]int
var dp_kill map[[4]int]int

func intervalIntersect(i1 [2]int, i2 [2]int) (bool){
	return !(i1[0]>=i2[1] || i2[0]>=i1[1])
}

func optimalCut(rects [][4]int, x []int, y []int, reg [4]int, seq [][6]int) ([][6]int, int){
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

		kill1 := 255
		kill2 := 255
		seq1 := make([][6]int, 1)
		seq2 := make([][6]int, 1)

		if len(rects1) < len(rects) && len(rects2) < len(rects) {
			seq1, kill1 = optimalCut(rects1, xx1, yy1, reg1, seq)
			seq2, kill2 = optimalCut(rects2, xx2, yy2, reg2, seq)
		}

		cuts[k] = kill1 + kill2 + kill_cur

		seq_cur := seq
		seq_cur = append(seq_cur, seq1...)
		seq_cur = append(seq_cur, seq2...)
		seqs[k] = seq_cur

		if kill_cur==0 {
			var cur [][6]int
			cur = append(cur, [6]int{reg[0], reg[1], reg[2], x[1+k], 0})
			seqf := append(cur, seq_cur...)
			dp_kill[reg] = cuts[k]
			dp_seq[reg] = seqf
			return seqf, cuts[k]
		}

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

		kill1 := 255
		kill2 := 255
		seq1 := make([][6]int, 1)
		seq2 := make([][6]int, 1)

		if len(rects1) < len(rects) && len(rects2) < len(rects) {
			seq1, kill1 = optimalCut(rects1, xx1, yy1, reg1, seq)
			seq2, kill2 = optimalCut(rects2, xx2, yy2, reg2, seq)
		}

		cuts[len(x) - 2 + k] = kill1 + kill2 + kill_cur

		seq_cur := seq
		seq_cur = append(seq_cur, seq1...)
		seq_cur = append(seq_cur, seq2...)
		seqs[len(x) - 2 + k] = seq

		if kill_cur==0 {
			var cur [][6]int
			cur = append(cur, [6]int{reg[0], reg[1], reg[2], reg[3], y[1+k], 1})
			seqf := append(cur, seq_cur...)
			dp_kill[reg] = cuts[k]
			dp_seq[reg] = seqf
			return seqf, cuts[k]
		}
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
	for k, rec1 := range rects {
		for _, rec2 := range rects[k+1:] {
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

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	var rects [][4]int
	for i:=0; i<n; i++ {
		var x1, x2, y1, y2 int
		fmt.Scanf("%d %d %d %d\n", &x1, &x2, &y1, &y2)
		rects = append(rects, [4]int{x1, x2, y1, y2})
	}
	//var rects = [][4]int{{3, 4, 2, 4}, {2, 4, 0, 2}, {2, 3, 2, 4}, {0, 2, 0, 4}}

	dp_seq = make(map[[4]int][][6]int)
	dp_kill = make(map[[4]int]int)

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

		fin_seq, killed := optimalCut(rects, x, y, reg, seq)
		fmt.Println(fin_seq)
		fmt.Println(killed)

	} else {
		fmt.Println("Invalid set!")
	}
}
