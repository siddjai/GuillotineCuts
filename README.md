# Guillotine Cuts
This repository contains all code written for my Bachelor's Thesis at IIIT-Delhi. Currently we are trying to enumerate optimal cut sequences for Axis Parallel Rectangles. The general problem dates back to the paper [Cutting Glass](https://dl.acm.org/citation.cfm?id=336223) [Pach et al] but the specific problem we're working on is presented directly in [On Guillotine Cutting Sequences](http://drops.dagstuhl.de/opus/volltexte/2015/5291/) [Abed et al]. In particular, we want to know if one can always save Omega(n) rectangles and exactly what fraction is feasible.

## Guide 

### 3 Phases 

- **Generating Tree** : In this phase we enumerate a set of permutations using a Generating Tree, explained in linked paper.
- **Permutation 2 Floorplan** : We interpret the permutation as a mosaic floorplan.
- **Compute Optimal-Cut-Seq**: Finally, we compute the optimal Guillotine cut sequence on the set of rectangles.

### Run & Reference

Each folder contains a *.go* code which was/will be actually used for enumeration and a *.py* code which was written as proof of concept to iron out implementation details.

### Possible Edits

- Use **uint8** in Generating Tree implementation, that could save memory if needed and speed up a bit.
- Instead of the *reg*, use an encoding of labels in ComputeOptimalCutSeq. This would add overhead in function but reduce size of *seq* and could potentially lead to speedup.

All points mentioned here need further analysis.
