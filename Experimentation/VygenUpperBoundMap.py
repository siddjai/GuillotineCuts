# Format for specifying a rectangle is (x1 x2 y1 y2)

from toposort import toposort, toposort_flatten

# Code snippet taken from RosettaCode : https://rosettacode.org/wiki/Topological_sort#Python
from functools import reduce

def toposort2(data):
    for k, v in data.items():
        v.discard(k) # Ignore self dependencies
    extra_items_in_deps = reduce(set.union, data.values()) - set(data.keys())
    data.update({item:set() for item in extra_items_in_deps})
    while True:
        ordered = set(item for item,dep in data.items() if not dep)
        if not ordered:
            break
        yield ' '.join(sorted(map(str, ordered)))
        data = {item: (dep - ordered) for item,dep in data.items()
                if item not in ordered}
    assert not data, "A cyclic dependency exists amongst %r" % data

# End of snippet

n = int(input())
recs = list()
for i in range(n):
	r = [int(k) for k in input().split()]
	r = tuple(r)
	recs.append(r)

N, E, W, S = set(), set(), set(), set()

for i in range(len(recs)):
	for j in range(i+1, len(recs)):
		r1 = recs[i]
		r2 = recs[j]
		t1 = (j+1, i+1)
		t2 = (i+1, j+1)
		if r1[1] <= r2[0]:
			E.add(t1)
			W.add(t2)
		if r1[0] >= r2[1]:
			E.add(t2)
			W.add(t1)
		if r1[3] <= r2[2]:
			N.add(t1)
			S.add(t2)
		if r1[2] >= r2[3]:
			N.add(t2)
			S.add(t1)

A1 = (S.difference(E)).union(W.difference(N))
A2 = (S.difference(W)).union(E.difference(N))
print(A1)
print(A2)
print()

l1 = dict()
for k in range(1, n+1):
	l1[k] = set()
for e in A1:
	l1[e[0]].add(e[1])

l2 = dict()
for k in range(1, n+1):
	l2[k] = set()
for e in A2:
	l2[e[0]].add(e[1])

ppi = list(toposort(l1))
prho = list(toposort(l2))
# Here you could try all linear extensions of this partial order
pi = list()
rho = list()
for s in ppi:
	pi = pi + list(s)
for s in prho:
	rho = rho + list(s)
print(*pi)
print(*rho)
print()

rhoo = list()
for k in rho:
	rhoo.append(pi[k-1])

print(*rhoo)
