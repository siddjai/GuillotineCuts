# k x n matrix
# 1 x m tiles
# m = k-1
# n = m + k - 1

import copy

po = set()

def makeOrder(k, mat):
	E = [[0 for i in range(k)] for j in range(k)]
	# If needed, add diagonal to make an actual partial order
	# for i in range(k): E[i][i] = 1
	for i in range(k):
		for j in range(k):
			if mat[i][1] >= mat[j][1]+(k-1): E[i][j] = 1

	ordstr = ""
	for i in range(k):
		ordstr += ''.join(str(e) for e in E[i])

	return ordstr

def makeTiling(k, mat):
	m = k-1
	n = 2*m
	tiling = [[0 for i in range(n)] for j in range(k)]
	for s in mat:
		for a in range(m):
			tiling[s[0]][s[1]+a] = 1

	return tiling

def reducible(k, mat):
	m = k-1
	for s in mat:
		if s[1] in range(m) and s[1]+m-1 in range(m,2*m):
			return False
	return True

def construct(k, mat, ind):
	for s in range(k):
		newmat = copy.copy(mat)
		newmat.append((ind, s))
		if ind == k-1:
			if not reducible(k, newmat):
				po.add(makeOrder(k, newmat))
				# tiling = makeTiling(k, newmat)
				# for i in range(k):
				# 	print(''.join(str(e) for e in tiling[i]))
				# print()

		else: construct(k, newmat, ind+1)


def countSets(k):
	po.clear()

	# Will store start of every tile added
	mat = list()

	construct(k, mat, 0)

	return len(po)

for k in range(3,10):
	print(k, countSets(k))