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
	x1, x2 = i1[0], i1[1]
	if (x1 > i2[0] and x1 < i2[1]) or (x2 > i2[0] and x2 < i2[1]):
		return True

	x1, x2 = i2[0], i2[1]
	if (x1 > i1[0] and x1 < i1[1]) or (x2 > i1[0] and x2 < i1[1]):
		return True

	return False

def optimalCut(rects, x, y, reg, seq):
	# rects : Rectangles in the current set
	# x : sorted list of X coordinates
	# y : sorted list of Y coordinates
	# reg : current bounded region
	# seq : sequence of cuts upto this set

	# RETURN
	# seq : seq of rectangles including this level
	# killed : no of rectangles killed including this level

	if len(rects)==1 or len(rects)==0:
		return seq, 0
	if len(rects)==2:
		rectsL = list(rects)
		rec1, rec2 = rectsL[0], rectsL[1]
		x1, x2, y1, y2 = (rec1[0], rec1[1]), (rec2[0], rec2[1]), (rec1[2], rec1[3]), (rec2[2], rec2[3])
		for xx in x[1:-1]:
			if (not intervalIntersect(x1, (xx,xx))) and (not intervalIntersect(x2, (xx,xx))):
				return seq + [(xx, 0)], 0
		for yy in y[1:-1]:
			if (not intervalIntersect(y1, (yy,yy))) and (not intervalIntersect(y2, (yy,yy))):
				return seq + [(yy, 1)], 0

	# Check if already memoized here
	key_rects = tuple(sorted(list(rects)))
	if key_rects in dp.keys():
		return dp[key_rects]

	m = len(x) + len(y) - 4
	cuts = [ 1000 for k in range(m) ]
	seqs = []
	for k in range(len(x) - 2):
		rects1 = set()
		boundary = False
		for rec in rects:
			if rec[0] < x[1+k]: rects1.add(rec)
			if rec[1] == x[1+k]: boundary = True

		rects2 = rects.difference(rects1)

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
		kill3 = 0
		for tup in rects:
			xi = tup[:2]
			if intervalIntersect(xi, (x[1+k], x[1+k])):
				kill3 += 1
		cuts[k] = kill1 + kill2 + kill3
		# Check if correct
		seqs.append(seq + seq1 + seq2)

	for k in range(len(y) - 2):
		rects1 = set()
		boundary = False
		for rec in rects:
			if rec[2] < y[1+k]: rects1.add(rec)
			if rec[3] == y[1+k]: boundary = True

		rects2 = rects.difference(rects1)

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
		kill3 = 0
		for tup in rects:
			yi = tup[2:]
			if intervalIntersect(yi, (y[1+k], y[1+k])):
				kill3 += 1
		cuts[len(x) - 2 + k] = kill1 + kill2 + kill3
		# Check if correct
		seqs.append(seq + seq1 + seq2)

	minPtr = 0

	for k in range(m):
		if cuts[k] < cuts[minPtr]: minPtr = k

	newLine = (1000, 0)
	if minPtr < len(x) - 1: newLine = (x[1 + minPtr], 0)
	else: newLine = (y[minPtr - len(x)], 1)

	# Add to dictionary here
	dp[key_rects] = ([newLine] + seqs[minPtr], cuts[minPtr])
	return [newLine] + seqs[minPtr], cuts[minPtr]


def sanityCheck(rects):
	for rec1 in rects:
		for rec2 in rects:
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
	reg = (0, 1000, 0, 1000)
	seq = []

	print(optimalCut(rects, x, y, reg, seq))

else: print("Rectangle set not valid")

