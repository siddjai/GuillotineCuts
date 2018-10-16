#!/usr/bin/env sage

def localExp(perm, a):
	# Local expansion as described in the paper
	newPerm = []
	for k in perm:
		if k < a: newPerm.append(k)
		else: newPerm.append(k+1)

	newPerm.append(a)

	return tuple(newPerm)

def isBaxter(perm):
	n = len(perm)

	# Memorise -41-
	steps = []
	for k in range(n-1):
		if perm[k] > perm[k+1] + 2: steps.append(k)

	# Avoid 2-41-3
	for s in steps:
		m, M = perm[s+1], perm[s]
		two, three = 1000, 0
		prefix, suffix = perm[:s], perm[s+2:]
		for k in prefix:
			if (k > m) and (k < M - 1):
				two = min(k, two)

		for k in suffix:
			if (k > two) and (k < M):
				three = k
				return False

	# Memorise -14-
	steps = []
	for k in range(n-1):
		if perm[k] < perm[k+1] - 2: steps.append(k)


	# Avoid 3-41-2
	for s in steps:
		m, M = perm[s], perm[s+1]
		two, three = 1000, 0
		prefix, suffix = perm[:s], perm[s+2:]
		for k in prefix:
			if (k > m + 1) and (k < M):
				three = max(k, three)

		for k in suffix:
			if (k < three) and (k > m):
				three = k
				return False

	return True


curLevel = set([(1,2,3), (1,3,2), (2,1,3), (3,1,2), (2,3,1), (3,2,1)])
level = 3
while level != 12:
	newLevel = set()
	for perm in curLevel:
		for a in range(1,level+2):
			newPerm = localExp(perm, a)
			if isBaxter(newPerm): newLevel.add(newPerm)

	gc1 = graph_classes.get_class("gc_15")
	c1 = 0
	for p in newLevel:
		n = len(p)
		d = dict()
		for i in range(n-1):
			p_i = p[i]
			l_i = []
			for j in range(i+1, n):
				if p[j] < p[i]:
					l_i.append(p[j])
			d[p_i - 1] = l_i
		g = Graph(d)
		if g not in gc1:
			c1 +=1
	print(c1)

	curLevel = newLevel
	level += 1
