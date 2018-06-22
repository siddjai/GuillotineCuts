# Guillotine Cuts
This repository contains all code written for my Bachelor's Thesis at IIIT-Delhi. Currently we are trying to enumerate optimal cut sequences for Axis Parallel Rectangles. The general problem dates back to the paper [Cutting Glass](https://dl.acm.org/citation.cfm?id=336223) [Pach et al] but the specific problem we're working on is presented directly in [On Guillotine Cutting Sequences](http://drops.dagstuhl.de/opus/volltexte/2015/5291/) [Abed et al]. In particular, we want to know if one can always save Omega(n) rectangles and exactly what fraction is feasible.

## Guide 

### /ref
This folder contains all the proof of concept code. This code will not be run directly for enumeration but it is simply for ironing out implementation details, and to help decide structure of more low level code.

### Possible Edits

- Use **uint8** in Generating Tree implementation, that could save memory if needed and speed up a bit.
- Instead of the *reg*, use an encoding of labels in ComputeOptimalCutSeq. This would add overhead in function but reduce size of *seq* and could potentially lead to speedup.

All points mentioned here need further analysis.
