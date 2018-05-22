# k x n matrix
# 1 x m tiles
# m = k-1
# n = m + 2k - 1

import copy

po = set()

def makeOrder(k, mat):
	# E = [[0 for i in range(k)] for j in range(k)]
	# # If needed, add diagonal to make an actual partial order
	# # for i in range(k): E[i][i] = 1
	# for i in range(k):
	# 	for j in range(k):
	# 		if mat[i][1] >= mat[j][1]+(k-1): E[i][j] = 1

	ordstr = ""
	for i in range(k):
		for j in range(k):
			if mat[i][1] >= mat[j][1]+mat[j][2]: ordstr+="1"
			else: ordstr+="0"

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
	# Seems to not be working
	m = k-1
	for i in range(2*k - 2):
		for s in mat:
			flag = True
			reason = False
			if s[1] <= i and s[1]+k-1 >= i:
				flag = False

			if s[1]+k-1 < i or s[1] > i:
				reason = True

		if reason and flag: return True
	return False

def construct(k, mat, ind):
	for s in range(k):
		newmat = copy.copy(mat)
		newmat.append((ind, s, k-1))
		if ind == k-1:
			if not reducible(k, newmat):
				# t = len(po)
				po.add(makeOrder(k, newmat))
				# if t < len(po):
				# 	tiling = makeTiling(k, newmat)
				# 	for i in range(k):
				# 		print(''.join(str(e) for e in tiling[i]))
				# 	print()

		else: construct(k, newmat, ind+1)

	for s in range(2*k-2):
		newmat = copy.copy(mat)
		newmat.append((ind, s, 1))
		if ind == k-1:
			if not reducible(k, newmat):
				# t = len(po)
				po.add(makeOrder(k, newmat))
				# if t < len(po):
				# 	tiling = makeTiling(k, newmat)
				# 	for i in range(k):
				# 		print(''.join(str(e) for e in tiling[i]))
				# 	print()
		else: construct(k, newmat, ind+1)

def countSets(k):
	po.clear()

	# Will store start of every tile added
	mat = list()

	construct(k, mat, 0)

	return len(po)

for k in range(3,10):
	print(k, countSets(k))
