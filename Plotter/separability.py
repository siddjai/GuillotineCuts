import math

def isNextTo(range1, range2):
	return range2[0] - range1[1] == 1 or range1[0] - range2[1] == 1

def combine(range1, range2):
	if range2[0] - range1[1] == 1:
		return (range1[0], range2[1])
	if range1[0] - range2[1] == 1:
		return (range2[0], range1[1])
	return None

def isSeparable(permstr):
	perm = tuple([int(k) for k in permstr.split()])
	stack = []
	for v in perm:
		range = (v, v)
		while len(stack) > 0 and isNextTo(range, stack[-1]):
			range = combine(range, stack.pop())
		stack.append(range)
	return len(stack) == 1
