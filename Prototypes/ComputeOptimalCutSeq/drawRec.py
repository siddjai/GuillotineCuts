def draw(rects, n):
	mat = [[0 for i in range(n)] for j in range(n)]
	for k in rects:
		rec = rects[k]
		for i in range(rec[0], rec[1]):
			for j in range(rec[2], rec[3]):
				mat[n - 1 - j][i] = k

	for k in range(n):
		print("".join(str(i) for i in mat[k]))

n = int(input())
rects = dict()
for k in range(n):
	this = [int(k) for k in input().split()]
	this = tuple(this)
	rects[k+1] = this

draw(rects, n)
