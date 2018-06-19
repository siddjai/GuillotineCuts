//// Format for specifying a rectangle:
//// (x1, x2, y1, y2) where x1 is the min and x2 is the max y coordinate.
//// Similarly for y
//
//// Format for specifying a line:
//// (a, b) where a is a number and b is a binary number where
//// 0 -> || to Y axis AND 1 -> || to X axis
//
//// Format for specifying an interval:
//// (a1, a2)
//
//package main
//
////func intervalIntersect(i1 [2]int, i2 [2]int) {
////	x1, x2 = i1[0], i1[1]
////	if (x1 > i2[0] and
////	x1 < i2[1]) || (x2 > i2[0] and x2 < i2[1]){
////return true
////}
////
////x1, x2 = i2[0], i2[1]
////if (x1 > i1[0] and x1 < i1[1]) || (x2 > i1[0] and x2 < i1[1]){
////return true
////}
////
////return false
////}
//
//func optimalCut() {
//	// rects : Rectangles in the current set
//	// x : sorted list of X coordinates
//	// y : sorted list of Y coordinates
//	// reg : current bounded region [REMOVED BECAUSE OF NO USE IN CODE]
//	// seq : sequence of cuts upto this set
//
//	// RETURN
//	// seq : seq of rectangles including this level
//	// killed : no of rectangles killed including this level
//}
//
//func sanityCheck(rects [][4]int):
//for _, rec1 := range rects:
//for _, rec2 := range rects:
//x1, x2, y1, y2 = (rec1[0], rec1[1]), (rec2[0], rec2[1]), (rec1[2], rec1[3]), (rec2[2], rec2[3])
//if intervalIntersect(x1, x2) && intervalIntersect(y1, y2) {
//return false
//}
//
//return true
//
//// Input in Go syntax
//n = int(input())
//
//rects = [][4]int
//for k = 0; k<n; k++ {
//
//this := []int
//// Do this but in Go syntax
//this = [int(k) for k in input().split()]
//rects = append(rects, this)
//}
//
//if sanityCheck(rects):
//x = []int
//y = []int
//for tup in rects:
//x = append(x, tup[0])
//x = append(x, tup[1])
//y = append(y, tup[2])
//y = append(y, tup[3])
//
//Sort(x)
//Sort(y)
//reg = [4]int{x[0], x[-1], y[0], y[-1]}
//seq = [][2]int
//
//print(optimalCut(rects, x, y, reg, seq)) else: print("Rectangle set not valid")
