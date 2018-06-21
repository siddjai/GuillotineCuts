// Given a Baxter permutation, this program constructs a corresponding floorplan
// Based on the mapping mentioned on page 15 in this thesis:
// https://www.cs.technion.ac.il/users/wwwb/cgi-bin/tr-get.cgi/2006/PHD/PHD-2006-11.pdf
// And the related paper
// Eyal Ackerman, Gill Barequet, and Ron Y. Pinter.  A bijection
// between permutations and floorplans, and its applications.
// Discrete Applied Mathematics, 154(12):1674–1684, 2006.

// Format for specifying a rectangle:
// (x1, x2, y1, y2) where x1 is the min and x2 is the max y coordinate.
// Similarly for y