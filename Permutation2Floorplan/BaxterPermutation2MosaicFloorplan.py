# Given a Baxter permutation, this program constructs a corresponding floorplan
# Based on the mapping mentioned on page 15 in this thesis:
# https://www.cs.technion.ac.il/users/wwwb/cgi-bin/tr-get.cgi/2006/PHD/PHD-2006-11.pdf
# And the related paper
# Eyal Ackerman, Gill Barequet, and Ron Y. Pinter.  A bijection
# between permutations and floorplans, and its applications.
# Discrete Applied Mathematics, 154(12):1674–1684, 2006.

# Format for specifying a rectangle:
# (x1, x2, y1, y2) where x1 is the min and x2 is the max y coordinate.
# Similarly for y

def draw(rects, n):
	rects = list(rects)
	mat = [[0 for i in range(n)] for j in range(n)]
	for k in range(len(rects)):
		rec = rects[k]
		for i in range(rec[0], rec[1]):
			for j in range(rec[2], rec[3]):
				mat[n - 1 - j][n - 1 - i] = k+1

	for k in range(n):
		print("".join(str(i) for i in mat[k]))



perm = [int(k) for k in input().split()]
perm = tuple(perm)
n = len(perm)

rects = dict()
rects[perm[0]] = (0, n, 0, n)
below = dict()
left = dict()
prevlabel = perm[0]
for k in range(1, n):
	if perm[k] < prevlabel:
		oldrect = rects[prevlabel]
		middle = (oldrect[2]+oldrect[3])//2

		rk = list(oldrect)
		rk[2] = middle
		rects[perm[k]] = tuple(rk)

		rpl = list(oldrect)
		rpl[3] = middle
		rects[prevlabel] = tuple(rpl)

		below[perm[k]] = prevlabel
		if prevlabel in left: left[perm[k]] = left[prevlabel]

		while perm[k] in left and left[perm[k]] > perm[k]:
			l = left[perm[k]]
			leftrect = rects[l]

			rp = list(rects[perm[k]])
			rp[0] = leftrect[0]
			rects[perm[k]] = tuple(rp)

			rl = list(rects[l])
			rl[3] = rp[2]
			rects[l] = tuple(rl)

			if l in left.keys(): left[perm[k]] = left[l]
			else: del left[perm[k]]

		prevlabel = perm[k]

	else:
		oldrect = rects[prevlabel]
		middle = (oldrect[0]+oldrect[1])//2

		rk = list(oldrect)
		rk[0] = middle
		rects[perm[k]] = tuple(rk)

		rpl = list(rects[prevlabel])
		rpl[1] = middle
		rects[prevlabel] = tuple(rpl)

		left[perm[k]] = prevlabel
		if prevlabel in below: below[perm[k]] = below[prevlabel] 
		
		while perm[k] in below and below[perm[k]] < perm[k]:
			b = below[perm[k]]
			belowrect = rects[b]

			rp = list(rects[perm[k]])
			rp[2] = belowrect[2]
			rects[perm[k]] = tuple(rp)

			rb = list(rects[b])
			rb[1] = rp[0]
			rects[b] = tuple(rb)

			if b in below.keys(): below[perm[k]] = below[b]
			else: del below[perm[k]]

		prevlabel = perm[k]

	curRects = set(rects.values())
	draw(curRects, n)
	print()

rects = set(rects.values())
print(rects)
draw(rects, n)