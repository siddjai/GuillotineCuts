# Thanos because you must kill half!

import matplotlib.pyplot as plt

plt.axes()
m, n = 20, 20
b = 1
l = m*b/2

for j in range(m):

	for k in range(n):
	## Do m and n need to be switched to represent rows and columns like our representation? IDK fix this if required pls
	## ALSO check the offsets. I took 2.5 == l + b/2 and 0.5 == b/2 by intuition, not derivation.

		# rectangle = plt.Rectangle((0.5 + k*(l+b/2) - (b/2)*j, 2.5 + (b/2)*k + (l+b/2)*j), l, b, fc='r') #1
		# plt.gca().add_patch(rectangle)
		rectangle = plt.Rectangle((b/2 + k*(l+b/2) - (b/2)*j, l + b/2 + (b/2)*k + (l+b/2)*j), l, b, fc='r') #1
		plt.gca().add_patch(rectangle)

		# rectangle = plt.Rectangle((2.5 + k*(l+b/2) - (b/2)*j, 1 + (b/2)*k + (l+b/2)*j), b, l, fc='r') #2
		rectangle = plt.Rectangle((l + b/2 + k*(l+b/2) - (b/2)*j, b + (b/2)*k + (l+b/2)*j), b, l, fc='r') #2
		plt.gca().add_patch(rectangle)

		# rectangle = plt.Rectangle((1 + k*(l+b/2) - (b/2)*j, 0 + (b/2)*k + (l+b/2)*j), l, b, fc='r') #3
		rectangle = plt.Rectangle((b + k*(l+b/2) - (b/2)*j, 0 + (b/2)*k + (l+b/2)*j), l, b, fc='r') #3
		plt.gca().add_patch(rectangle)

		# rectangle = plt.Rectangle((0 + k*(l+b/2) - (b/2)*j, 0.5 + (b/2)*k + (l+b/2)*j), b, l, fc='r') #4
		rectangle = plt.Rectangle((0 + k*(l+b/2) - (b/2)*j, b/2 + (b/2)*k + (l+b/2)*j), b, l, fc='r') #4
		plt.gca().add_patch(rectangle)

		rectangle = plt.Rectangle((b + k*(l+b/2) - (b/2)*j, l/3 + 5*b/6 + (b/2)*k + (l+b/2)*j), l-b/2, (l-b/2)/3, fc='b') #Inner
		plt.gca().add_patch(rectangle)

# Use this to play around with cuts, you can infer the syntax
line = plt.Line2D((0, 0), (180, 0), lw=1)
plt.gca().add_line(line)


plt.axis('scaled')
plt.show()
# plt.savefig('3,3grid.png', bbox_inches='tight')
