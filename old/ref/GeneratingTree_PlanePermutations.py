# Plane permutations are those which avoid the pattern 213'54
# Avoiding 213'54 is equivalent to avoiding 2-14-3
# This generating tree based enumeration is based on the following paper:
# https://arxiv.org/pdf/1702.04529.pdf

def localExp(perm, a):
	# Local expansion as described in the paper
	newPerm = []
	for k in perm:
		if k < a: newPerm.append(k)
		else: newPerm.append(k+1)

	newPerm.append(a)

	return tuple(newPerm)

def isPlane(perm):
	n = len(perm)
	steps = []
	for k in range(n-1):
		if perm[k] < perm[k+1] - 1: steps.append(k)

	for s in steps:
		m, M = perm[s], perm[s+1]
		two, three = 1000, 0
		prefix, suffix = perm[:s], perm[s+2:]
		for k in prefix:
			if (k > m) and (k < M - 1):
				two = min(k, two)

		for k in suffix:
			if (k > two) and (k < M):
				three = k
				return False

	return True


curLevel = set([(1,2,3), (1,3,2), (2,1,3), (3,1,2), (2,3,1), (3,2,1)])
level = 3
while level != 20:
	newLevel = set()
	for perm in curLevel:
		for a in range(1,level+2):
			newPerm = localExp(perm, a)
			if isPlane(newPerm): newLevel.add(newPerm)

	print(len(newLevel))
	curLevel = newLevel
	level += 1
