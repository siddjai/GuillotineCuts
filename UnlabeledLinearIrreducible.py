# k x n matrix
# 1 x m & 1 x 1 tiles 
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
	n = 2*(k-1)
	tiling = [[0 for i in range(n)] for j in range(k)]
	for s in mat:
		for a in range(s[2]):
			tiling[s[0]][s[1]+a] = 1

	return tiling

def reducible(k, mat):
	# Bottleneck
	for i in range(1, 2*k - 2):
		flag = True
		rleft = False
		right = False
		for s in mat:
			if s[1] < i and s[1]+s[2] > i:
				flag = False
				# print(i, "f")

			if s[1]+s[2] <= i:
				rleft = True
				# print(i, "rl")

			if s[1] >= i:
				right = True
				# print(i, "rr")

		if rleft and right and flag: return True

	return False

def addWithTiling(k, newmat):
	t = len(po)
	po.add(makeOrder(k, newmat))
	if t < len(po):
		tiling = makeTiling(k, newmat)
		for i in range(k):
			print(''.join(str(e) for e in tiling[i]))
		print()

def construct(k, mat, ind):
	for s in range(k):
		newmat = copy.copy(mat)
		newmat.append((ind, s, k-1))
		if ind == k-1:
			if not reducible(k, newmat):
				# addWithTiling(k, newmat)
				po.add(makeOrder(k, newmat))

		else: construct(k, newmat, ind+1)

	for s in range(2*k-2):
		newmat = copy.copy(mat)
		newmat.append((ind, s, 1))
		if ind == k-1:
			if not reducible(k, newmat):
				# addWithTiling(k, newmat)
				po.add(makeOrder(k, newmat))

		else: construct(k, newmat, ind+1)

	# for s in range(2*k-3):
	# 	newmat = copy.copy(mat)
	# 	newmat.append((ind, s, 2))
	# 	if ind == k-1:
	# 		if not reducible(k, newmat):
	# 			# addWithTiling(k, newmat)
	# 			po.add(makeOrder(k, newmat))

	# 	else: construct(k, newmat, ind+1)

	# for s in range(2*k-4):
	# 	newmat = copy.copy(mat)
	# 	newmat.append((ind, s, 3))
	# 	if ind == k-1:
	# 		if not reducible(k, newmat):
	# 			# addWithTiling(k, newmat)
	# 			po.add(makeOrder(k, newmat))

	# 	else: construct(k, newmat, ind+1)

def countSets(k):
	po.clear()

	# Will store start of every tile added
	mat = list()

	construct(k, mat, 0)

	return len(po)

# sim = [(0,0,2), (1,2,2), (2,2,2)]
# print(reducible(3, sim))

# print(4, countSets(4))

for k in range(4, 6):
	print(k, countSets(k))