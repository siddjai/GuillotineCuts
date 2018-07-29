from itertools import combinations

def reduce(perm):
	s = list(sorted(perm))
	for i in range(len(s)):
		perm = [i+1 if x==s[i] else x for x in perm]
	return perm


count = 0
violations = list()
f1 = [2, 4, 1, 3]
f2 = [3, 1, 4, 2]
perm = [int(k) for k in input().split()]
for c in combinations(perm, 4):
	r = reduce(c)
	if r==f1 or r==f2:
		count+=1
		violations.append([c, r])

print(count)
for v in violations:
	print(v[0], v[1])