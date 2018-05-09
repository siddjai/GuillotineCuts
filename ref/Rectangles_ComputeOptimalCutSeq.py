# Format for specifying a rectangle:
# (x1, x2, y1, y2) where x1 is the min and x2 is the max y coordinate.
# Similarly for y

# Format for specifying a line:
# (a, b) where a is a number and b is a binary number where
# 0 -> || to Y axis AND 1 -> || to X axis (

# Format for specifying an interval:
# (a1, a2)

def intervalIntersect(i1, i2):
	x1, x2 = i1[0], i1[1]
	if (x1 > i2[0] and x1 < i2[1]) or (x2 > i2[0] and x2 < i2[1]):
		return True
	return False

def optimalCut(rects, x, y, reg, seq, killed):
	# rects : Rectangles in the current set
	# x : sorted list of X coordinates
	# y : sorted list of Y coordinates
	# reg : current bounded region
	# seq : sequence of cuts upto this set
	# killed: no of rectangles killed upto this set

	# RETURN
	# seq : seq of rectangles including this level
	# killed : no of rectangles killed including this level

	if len(rects)==1:
		return seq, killed
	elif len(rects)==2:
		i1 = (rects[0][0], rects[0][1])
		i2 = (rects[1][0], rects[1][1])
		if not intervalIntersect(i1, i2):
			if rects[0][1] < rects[1][0]:
				return seq + [(rects[0][1], 0)], killed
			else:
				return seq + [(rects[1][1], 0)], killed
		else:
			if rects[0][3] < rects[1][2]:
				return seq + [(rects[0][3], 1)], killed
			else:
				return seq + [(rects[1][3], 1)], killed

	m = len(x) + len(y) - 2
	cuts = [ 1000 for k in range(m) ]
	seqs = []
	for k in range(len(x) - 1):
		# Compute all parameters ^
		seq1, kill1 = optimalCut(rects1, xx1, yy1, reg1, seq, killed)
		seq2, kill2 = optimalCut(rects2, xx2, yy2, reg2, seq, killed)
		cuts[k] = kill1 + kill2
		# Check if correct
		seqs.append(seq + seq1 + seq2)

	for k in range(len(y) - 1):
		# Compute all parameters ^
		seq1, kill1 = optimalCut(rects1, xx1, yy1, reg1, seq, killed)
		seq2, kill2 = optimalCut(rects2, xx2, yy2, reg2, seq, killed)
		cuts[len(x) - 1 + k] = kill1 + kill2
		# Check if correct
		seqs.append(seq + seq1 + seq2)

	minPtr = 0

	for k in range(m):
		if cuts[k] < cuts[minPtr]: minPtr = k

	newLine = (1000, 0)
	if minPtr < len(x) - 1: newLine = (x[1 + minPtr], 0)
	else: newLine = (y[minPtr - len(x)], 1)

	return newLine + seqs[minPtr], killed + cuts[minPtr]

n = int(input())
rects = []
for k in range(n):
	this = [int(k) for k in input().split()]
	this = tuple(this)
	rects.append(this)

x = set()
y = set()
for tup in rects:
	x.add(tup[0])
	x.add(tup[1])
	y.add(tup[2])
	y.add(tup[3])

x = sorted(list(x))
y = sorted(list(y))
reg = (1000, 1000, 1000, 1000)
seq = []
killed = 0

optimalCut(rects, x, y, reg, seq, killed)

