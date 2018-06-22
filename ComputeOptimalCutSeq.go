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
//	"bufio"
	"fmt"
	"sort"
)

// // Taken from StackOverflow
// // https://stackoverflow.com/questions/23192262/how-would-you-set-and-clear-a-single-bit-in-go
// func setBit(n int, pos uint) int {
//     n |= (1 << pos)
//     return n
// }

// func clearBit(n int, pos uint) int {
//     return n &^ (1 << pos)
// }

// func hasBit(n int, pos uint) bool {
//     val := n & (1 << pos)
//     return (val > 0)
// }

// // --- end of code snippet

// func encode(labels []int) {
// 	// Expectation: labels will be <= 20
// 	// Therefore resulting number can be stored in <int>
// 	e := 0
// 	for _, l := range labels {
// 		e = setBit(e, l)
// 	}

// 	return e
// }

var dp_seq map[[4]int][][6]int
var dp_kill map[[4]int]int

// var reader *bufio.Reader = bufio.NewReader(os.Stdin)
// func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func intervalIntersect(i1 [2]int, i2 [2]int) (bool){
	x1 := i1[0]
	x2 := i1[1]
	if (x1 > i2[0] && x1 < i2[1]) || (x2 > i2[0] && x2 < i2[1]){
		return true
	}

	x1 = i2[0]
	x2 = i2[1]
	if (x1 > i1[0] && x1 < i1[1]) || (x2 > i1[0] && x2 < i1[1]){
		return true
	}

	return false
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

	killed, ok := dp_kill[reg]
	if ok {
		sseq := dp_seq[reg]
		return sseq, killed
	}

	m := len(x) + len(y) - 4
	cuts := make([]int, m)
	seqs := make(map[int][][6]int)

	for i:=0; i<m; i++ {
		//Arbritary constant
		cuts[i] = 1000
	}

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
			} else if rec[0] < x[1+k] {
				rects1 = append(rects1, rec)
			} else {
				rects2 = append(rects2, rec)
			}

			if rec[1] == x[1+k] {boundary = true}

		}

		xx1 := x[:2+k]
		xx2 := x[2+k:]
		if boundary { xx2 = append([]int{x[1+k]}, xx2...) }

		var yy1 []int
		for _, rec := range rects1 {
			yy1 = append(yy1, rec[2])
			yy1 = append(yy1, rec[3])
		}

		reg1 := reg 
		reg1[1] = x[1+k]

		var yy2 []int
		for _, rec := range rects2 {
			yy2 = append(yy2, rec[2])
			yy2 = append(yy2, rec[3])
		}

		reg2 := reg 
		reg2[0] = x[1+k]

		sort.Slice(yy1, func(i, j int) bool { return i<j })
		sort.Slice(yy2, func(i, j int) bool { return i<j })

		seq1, kill1 := optimalCut(rects1, xx1, yy1, reg1, seq)
		seq2, kill2 := optimalCut(rects2, xx2, yy2, reg2, seq)

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
			} else if rec[2] < y[1+k] {
				rects1 = append(rects1, rec)
			} else {
				rects2 = append(rects2, rec)
			}

			if rec[3] == y[1+k] {boundary = true}
		}

		yy1 := x[:2+k]
		yy2 := x[2+k:]
		if boundary { yy2 = append([]int{y[1+k]}, yy2...) }

		var xx1 []int
		for _, rec := range rects1 {
			xx1 = append(xx1, rec[0])
			xx1 = append(xx1, rec[1])
		}

		reg1 := reg 
		reg1[3] = y[1+k]

		var xx2 []int 
		for _, rec := range rects2 {
			xx2 = append(xx2, rec[0])
			xx2 = append(xx2, rec[1])
		}

		reg2 := reg 
		reg2[2] = y[1+k]

		sort.Slice(xx1, func(i, j int) bool { return i<j })
		sort.Slice(xx2, func(i, j int) bool { return i<j })

		seq1, kill1 := optimalCut(rects1, xx1, yy1, reg1, seq)
		seq2, kill2 := optimalCut(rects2, xx2, yy2, reg2, seq)
	
		cuts[len(x) - 2 + k] = kill1 + kill2 + kill_cur

		seq = append(seq, seq1...)
		seq = append(seq, seq2...)
		seqs[len(x) - 2 + k] = seq
	}

	minPtr := 0
	for k:=0; k<m; k++ {
		if cuts[k] < cuts[minPtr] {minPtr = k}
	}

	newLine := [2]int{1000, 0}
	if minPtr < len(x) - 2 {
		newLine = [2]int{x[1+minPtr], 0}
	} else {
		newLine = [2]int{y[minPtr - len(x)], 1}
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

func main() {
	// var n int
	// fmt.Scanf("%d\n", &n)
	// var rects [][4]int
	// for i:=0; i<n; i++ {
	// 	var x1, x2, y1, y2 int
	// 	fmt.Scanf("%d %d %d %d\n", &x1, &x2, &y1, &y2)
	// 	rects = append(rects, [4]int{x1, x2, y1, y2})
	// }
	var rects = [][4]int{{3, 4, 2, 4}, {2, 4, 0, 2}, {2, 3, 2, 4}, {0, 2, 0, 4}}

	dp_seq = make(map[[4]int][][6]int)
	dp_kill = make(map[[4]int]int)

	if sanityCheck(rects) {
		var x []int
		var y []int
		for _, tup := range rects {
			x = append(x, tup[0])
			x = append(x, tup[1])
			y = append(y, tup[2])
			y = append(y, tup[3])
		}

		sort.Slice(x, func(i, j int) bool { return i<j })
		sort.Slice(y, func(i, j int) bool { return i<j })
		reg := [4]int{x[0], x[len(x)-1], y[0], y[len(x)-1]}
		var seq [][6]int

		fin_seq, killed := optimalCut(rects, x, y, reg, seq)
		fmt.Println(fin_seq)
		fmt.Println(killed)

	} else {
		fmt.Println("Invalid set!")
	}
}
