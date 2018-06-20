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

perm = [int(k) for k in input.split()]
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
		middle = (oldrect[2]+oldrect[3])/2
		rects[perm[k]] = (oldrect[0], oldrect[1], middle, oldrect[3])
		rects[prevlabel] = (oldrect[0], oldrect[1], oldrect[2], middle)
		below[perm[k]] = prevlabel
		# Implement while
		prevlabel = perm[k]
	else:
		oldrect = rects[prevlabel]
		middle = (oldrect[0]+oldrect[1])/2
		rects[perm[k]] = (middle, oldrect[1], oldrect[2], oldrect[3])
		rects[prevlabel] = (oldrect[0], middle, oldrect[2], oldrect[3])
		left[perm[k]] = prevlabel
		# Implement while
		prevlabel = perm[k]

rects = set(rects.values())
