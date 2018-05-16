# k x n matrix
# 1 x m tiles

import copy

po = set()

def makeOrder(k, mat):
	E = [[0 for i in range(k)] for j in range(k)]
	for i in range(k): E[i][i] = 1
	for i in range(k):
		for j in range(k):
			if mat[i][1] >= mat[j][1]+(k-1): E[i][j] = 1

	ordstr = ""
	for i in range(k):
		ordstr += ''.join(str(e) for e in E[i])

	return ordstr

def construct(k, m, n, mat, ind, alleft):
	if alleft:
		for s in range(m):
			newmat = copy.deepcopy(mat)
			newmat.append((ind, s))
			if ind == k-1: po.add(makeOrder(k, newmat))
			else:
				if s==0: construct(k, m, n, newmat, ind+1, True)
				else: construct(k, m, n, newmat, ind+1, False)

	else:
		for s in range(m+1):
			newmat = copy.deepcopy(mat)
			newmat.append((ind, s))
			if ind == k-1: po.add(makeOrder(k, newmat))
			else: construct(k, m, n, newmat, ind+1, False)


def countSets(k):
	po.clear()
	m = k-1
	n = m + k - 1
	# Will store start of every tile added
	mat = list()

	construct(k, m, n, mat, 0, False)
	return len(po)


for k in range(3,10):
	print(k, countSets(k))