# Format for specifying a rectangle:
# (x1, x2, y1, y2) where x1 is the min and x2 is the max y coordinate.
# Similarly for y

# Format for specifying a line:
# (a, b) where a is a number and b is a binary number where
# 0 -> || to Y axis AND 1 -> || to X axis

# Format for specifying an interval:
# (a1, a2)

dp = dict()

def intervalIntersect(i1, i2):
	return not (i1[0]>=i2[1] or i2[0]>=i1[1])

def optimalCut(rects, x, y, reg, seq):
	# rects : Rectangles in the current set
	# x : sorted list of X coordinates
	# y : sorted list of Y coordinates
	# reg : current bounded region
	# seq : sequence of cuts upto this set

	# RETURN
	# seq : seq of rectangles including this level
	# killed : no of rectangles killed including this level

	if len(rects)<=3:
		return seq, 0

	# Check if already memoized here
	if reg in dp:
		return dp[reg]

	# Sufficient to try all boundaries of rectangles
	m = len(x) + len(y) - 4
	cuts = [ 1000 for k in range(m) ]
	seqs = []
	for k in range(len(x) - 2):
		rects1, rects2 = set(), set()
		boundary = False
		kill_cur = 0
		for rec in rects:
			xi = rec[:2]
			if intervalIntersect(xi, (x[1+k], x[1+k])):
				kill_cur += 1
			elif rec[1] <= x[1+k]: rects1.add(rec)
			else: rects2.add(rec)

			if rec[0] == x[1+k]: boundary = True

		xx1 = x[:2+k]
		xx2 = x[2+k:]
		if boundary: xx2 = [x[1+k]] + xx2

		yy1 = set()
		for tup in rects1:
			yy1.add(tup[2])
			yy1.add(tup[3])

		reg1 = list(reg)
		reg1[1] = x[1+k] 
		reg1 = tuple(reg1)

		yy2 = set()
		for tup in rects2:
			yy2.add(tup[2])
			yy2.add(tup[3])

		yy1 = sorted(list(yy1))
		yy2 = sorted(list(yy2))

		reg2 = list(reg) 
		reg2[0] = x[1+k]
		reg2 = tuple(reg2)

		seq1, kill1 = optimalCut(rects1, xx1, yy1, reg1, seq)
		seq2, kill2 = optimalCut(rects2, xx2, yy2, reg2, seq)

		cuts[k] = kill1 + kill2 + kill_cur

		seqs.append(seq + seq1 + seq2)

	for k in range(len(y) - 2):
		rects1, rects2 = set(), set()
		boundary = False
		kill_cur = 0
		for rec in rects:
			yi = rec[2:]
			if intervalIntersect(yi, (y[1+k], y[1+k])):
				kill_cur += 1
			elif rec[3] <= y[1+k]: rects1.add(rec)
			else: rects2.add(rec)

			if rec[2] == y[1+k]: boundary = True

		yy1 = x[:2+k]
		yy2 = x[2+k:]
		if boundary: yy2 = [y[1+k]] + yy2

		xx1 = set()
		for tup in rects1:
			xx1.add(tup[0])
			xx1.add(tup[1])

		reg1 = list(reg)
		reg1[3] = y[1+k]
		reg1 = tuple(reg1) 

		xx2 = set()
		for tup in rects2:
			xx2.add(tup[0])
			xx2.add(tup[1])

		xx1 = sorted(list(xx1))
		xx2 = sorted(list(xx2))

		reg2 = list(reg) 
		reg2[2] = y[1+k]
		reg2 = tuple(reg2)

		seq1, kill1 = optimalCut(rects1, xx1, yy1, reg1, seq)
		seq2, kill2 = optimalCut(rects2, xx2, yy2, reg2, seq)

		cuts[len(x) - 2 + k] = kill1 + kill2 + kill_cur

		seqs.append(seq + seq1 + seq2)

	minPtr = 0

	for k in range(m):
		if cuts[k] < cuts[minPtr]: minPtr = k

	newLine = (1000, 0)
	if minPtr < len(x) - 2: newLine = (x[1 + minPtr], 0)
	else: newLine = (y[minPtr - len(x) - 2], 1)

	# Add to dictionary here
	dp[reg] = ([reg, newLine] + seqs[minPtr], cuts[minPtr])
	return [[reg,newLine]] + seqs[minPtr], cuts[minPtr]


def sanityCheck(rects):
	rects = list(rects)
	n = len(rects)
	for rec1 in rects[:n-1]:
		for rec2 in rects[k+1:]:
			x1, x2, y1, y2 = (rec1[0], rec1[1]), (rec2[0], rec2[1]), (rec1[2], rec1[3]), (rec2[2], rec2[3])
			if intervalIntersect(x1, x2) and intervalIntersect(y1, y2): return False

	return True 

n = int(input())
rects = set()
for k in range(n):
	this = [int(k) for k in input().split()]
	this = tuple(this)
	rects.add(this)

if sanityCheck(rects):
	x = set()
	y = set()
	for tup in rects:
		x.add(tup[0])
		x.add(tup[1])
		y.add(tup[2])
		y.add(tup[3])

	x = sorted(list(x))
	y = sorted(list(y))
	reg = (x[0], x[-1], y[0], y[-1])
	seq = []

	print(optimalCut(rects, x, y, reg, seq))

else: print("Rectangle set not valid")

