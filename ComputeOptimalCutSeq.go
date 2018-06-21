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

func optimalCut() {
	// rects : Rectangles in the current set
	// x : sorted list of X coordinates
	// y : sorted list of Y coordinates
	// reg : current bounded region [REMOVED BECAUSE OF NO USE IN CODE]
	// seq : sequence of cuts upto this set

	// RETURN
	// seq : seq of rectangles including this level
	// killed : no of rectangles killed including this level
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
	// Input in Go syntax
	var n int
	fmt.Scanf("%d\n", &n)
	var rects [][4]int
	for i:=0; i<n; i++ {
		var x1, x2, y1, y2 int
		fmt.Scanf("%d %d %d %d\n", &x1, &x2, &y1, &y2)
		rects = append(rects, [4]int{x1, x2, y1, y2})
	}
	//var rects = [][4]int{{3, 4, 2, 4}, {2, 4, 0, 2}, {2, 3, 2, 4}, {0, 2, 0, 4}}

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
		// reg := [4]int{x[0], x[len(x)-1], y[0], y[len(y)-1]}
		// var seq [][2]int

		fmt.Println("Init\n")
	} else {
		fmt.Println("Invalid set!")
	}
}
