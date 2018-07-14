package mocs

import (
	"GuillotineCuts/BP2FP"
	"GuillotineCuts/div"
	"GuillotineCuts/eqv"
	"GuillotineCuts/pkg"
	"GuillotineCuts/sep"
	"bytes"
)

func getCuts(perm []uint8, rects [][4]uint8) (map[uint8]bool, map[uint8]bool) {
	hcuts := make(map[uint8]bool)
	vcuts := make(map[uint8]bool)
	n := uint8(len(perm))
	for _, rect := range rects {
		vcuts[rect[0]] = true
		vcuts[rect[1]] = true
		hcuts[rect[2]] = true
		hcuts[rect[3]] = true
	}
	delete(vcuts, 0)
	delete(vcuts, n)
	delete(hcuts, 0)
	delete(hcuts, n)
	eqvs := eqv.GetEqv(perm)
	mid := (n + 1) / 2
	if bytes.Equal(perm, eqvs[6]) {
		for cut, _ := range vcuts {
			if cut > mid {
				delete(vcuts, cut)
			}
		}
		return vcuts, make(map[uint8]bool)
	} else if bytes.Equal(perm, eqvs[3]) {
		for cut, _ := range vcuts {
			if cut > mid {
				delete(vcuts, cut)
			}
		}
		for cut, _ := range hcuts {
			if cut > mid {
				delete(hcuts, cut)
			}
		}
		return vcuts, hcuts
	} else if bytes.Equal(perm, eqvs[1]) || bytes.Equal(perm, eqvs[2]) {
		return vcuts, make(map[uint8]bool)
	} else if bytes.Equal(perm, eqvs[4]) {
		for cut, _ := range hcuts {
			if cut > mid {
				delete(hcuts, cut)
			}
		}
		return vcuts, hcuts
	} else if bytes.Equal(perm, eqvs[7]) {
		for cut, _ := range vcuts {
			if cut > mid {
				delete(vcuts, cut)
			}
		}
		return vcuts, hcuts
	} else {
		return vcuts, hcuts
	}
}

func finalMosaicOCS(m *Manager, worker_id uint16, perm []uint8) uint8 {
	signature := pkg.ToString(perm)
	cache_id := len(perm) - 5
	m.cacheLocks[cache_id].Lock()
	// Cache not existed
	if _, ok := m.caches[cache_id][signature]; !ok {
		m.addSignature(perm, 0, worker_id) // Temporarily calculated by this worker_id
		m.cacheLocks[cache_id].Unlock()

		//Divisible
		if cut := div.IsMosaicDivisible(perm); cut > 0 {
			return mosaicOCSUtil(m, worker_id, pkg.ToPermutation(perm[:cut])) + mosaicOCSUtil(m, worker_id, pkg.ToPermutation(perm[cut:]))
		}

		// Non-divisible
		minCost := m.n
		rects := BP2FP.BP2FP(perm)
		vcuts, hcuts := getCuts(perm, rects)
		for cut, _ := range vcuts {
			left := make([]uint8, 0)
			right := make([]uint8, 0)
			for _, p := range perm {
				rect := rects[p-1]
				if rect[0] >= cut {
					right = append(right, p)
				}
				if rect[1] <= cut {
					left = append(left, p)
				}
			}
			cost := uint8(len(perm) - len(left) - len(right))
			cost += mosaicOCSUtil(m, worker_id, pkg.ToPermutation(left)) + mosaicOCSUtil(m, worker_id, pkg.ToPermutation(right))
			minCost = pkg.Min(minCost, cost)
		}
		for cut, _ := range hcuts {
			left := make([]uint8, 0)
			right := make([]uint8, 0)
			for _, p := range perm {
				rect := rects[p-1]
				if rect[2] >= cut {
					right = append(right, p)
				}
				if rect[3] <= cut {
					left = append(left, p)
				}
			}
			cost := uint8(len(perm) - len(left) - len(right))
			cost += mosaicOCSUtil(m, worker_id, pkg.ToPermutation(left)) + mosaicOCSUtil(m, worker_id, pkg.ToPermutation(right))
			minCost = pkg.Min(minCost, cost)
		}
		return minCost
	} else { // Cache existed
		m.cacheLocks[cache_id].Unlock()
		return 0 // fake return since there is someone else doing the job
	}
}

func cacheMosaicOCS(m *Manager, worker_id uint16, perm []uint8) {
	signature := pkg.ToString(perm)
	cache_id := len(perm) - 5
	m.cacheLocks[cache_id].Lock()
	// Cache not existed
	if _, ok := m.caches[cache_id][signature]; !ok {
		m.addSignature(perm, 0, worker_id) // Temporarily calculated by this worker_id
		m.workerLocks[worker_id].Lock()
		m.cacheLocks[cache_id].Unlock()
		// Start calculating
		minCost := m.n
		rects := BP2FP.BP2FP(perm)
		vcuts, hcuts := getCuts(perm, rects)
		for cut, _ := range vcuts {
			left := make([]uint8, 0)
			right := make([]uint8, 0)
			for _, p := range perm {
				rect := rects[p-1]
				if rect[0] >= cut {
					right = append(right, p)
				}
				if rect[1] <= cut {
					left = append(left, p)
				}
			}
			cost := uint8(len(perm) - len(left) - len(right))
			cost += mosaicOCSUtil(m, worker_id, pkg.ToPermutation(left)) + mosaicOCSUtil(m, worker_id, pkg.ToPermutation(right))
			minCost = pkg.Min(minCost, cost)
		}
		for cut, _ := range hcuts {
			left := make([]uint8, 0)
			right := make([]uint8, 0)
			for _, p := range perm {
				rect := rects[p-1]
				if rect[2] >= cut {
					right = append(right, p)
				}
				if rect[3] <= cut {
					left = append(left, p)
				}
			}
			cost := uint8(len(perm) - len(left) - len(right))
			cost += mosaicOCSUtil(m, worker_id, pkg.ToPermutation(left)) + mosaicOCSUtil(m, worker_id, pkg.ToPermutation(right))
			minCost = pkg.Min(minCost, cost)
		}
		m.cacheLocks[cache_id].Lock()
		m.addSignature(perm, minCost, worker_id)
		m.cacheLocks[cache_id].Unlock()
		m.workerLocks[worker_id].Unlock()
	} else { // Cache existed then move on, do nothing
		m.cacheLocks[cache_id].Unlock()
	}
}

// Assume worker_id has locked before the call of this function
func mosaicOCSUtil(m *Manager, worker_id uint16, perm []uint8) uint8 {
	// Separable mosaic
	// O(n)
	if sep.IsSeparable(perm) {
		return 0
	}

	// Divisible mosaic
	// O(nlogn)
	if cut := div.IsMosaicDivisible(perm); cut > 0 {
		return mosaicOCSUtil(m, worker_id, pkg.ToPermutation(perm[:cut])) + mosaicOCSUtil(m, worker_id, pkg.ToPermutation(perm[cut:]))
	}

	// Non-divisible mosaic
	// O(nlogn)
	// Need to cache
	signature := pkg.ToString(perm)
	cache_id := len(perm) - 5
	m.cacheLocks[cache_id].Lock()
	if cache, ok := m.caches[cache_id][signature]; ok {
		m.cacheLocks[cache_id].Unlock()
		if cache.cost == 0 {
			m.workerLocks[cache.worker_id].Lock()
			m.workerLocks[cache.worker_id].Unlock()
			return cache.cost
		}
		return cache.cost
	}
	m.addSignature(perm, 0, worker_id)
	m.cacheLocks[cache_id].Unlock()
	minCost := m.n
	rects := BP2FP.BP2FP(perm)
	vcuts, hcuts := getCuts(perm, rects)

	for cut, _ := range vcuts {
		left := make([]uint8, 0)
		right := make([]uint8, 0)
		for _, p := range perm {
			rect := rects[p-1]
			if rect[0] >= cut {
				right = append(right, p)
			}
			if rect[1] <= cut {
				left = append(left, p)
			}
		}
		cost := uint8(len(perm) - len(left) - len(right))
		cost += mosaicOCSUtil(m, worker_id, pkg.ToPermutation(left)) + mosaicOCSUtil(m, worker_id, pkg.ToPermutation(right))
		minCost = pkg.Min(minCost, cost)
	}

	for cut, _ := range hcuts {
		left := make([]uint8, 0)
		right := make([]uint8, 0)
		for _, p := range perm {
			rect := rects[p-1]
			if rect[2] >= cut {
				right = append(right, p)
			}
			if rect[3] <= cut {
				left = append(left, p)
			}
		}
		cost := uint8(len(perm) - len(left) - len(right))
		cost += mosaicOCSUtil(m, worker_id, pkg.ToPermutation(left)) + mosaicOCSUtil(m, worker_id, pkg.ToPermutation(right))
		minCost = pkg.Min(minCost, cost)
	}

	m.cacheLocks[cache_id].Lock()
	m.addSignature(perm, minCost, worker_id)
	m.cacheLocks[cache_id].Unlock()
	return minCost
}
