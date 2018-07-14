package mocs

import (
	"GuillotineCuts/bxt"
	"GuillotineCuts/div"
	"GuillotineCuts/eqv"
	"GuillotineCuts/pkg"
	"GuillotineCuts/sep"
	"fmt"
	"os"
	"sync"
	"time"
)

type Manager struct {
	n, maxCost              uint8
	maxPerms                [][]uint8
	workersNumber           uint16
	workerLocks, cacheLocks []sync.Mutex
	maxLock, countLock      sync.Mutex
	jobs                    chan []uint8
	wgWorkers               sync.WaitGroup
	caches                  []map[string]Cache
	baxter_numbers          [17]uint64
	count                   uint64
	isCached                bool
}

type Cache struct {
	cost      uint8 // 0 if it is currently calculating by worker_id
	worker_id uint16
}

type Result struct {
	maxCost  uint8
	maxPerms [][]uint8
}

func NewManager(n uint8, workersNumber uint16, maxJobs uint32, isCached bool) *Manager {
	m := Manager{n: n, workersNumber: workersNumber, isCached: isCached}
	m.baxter_numbers = [17]uint64{0, 1, 2, 6, 22, 92, 422, 2074, 10754, 58202, 326240, 1882960, 11140560, 67329992, 414499438, 2593341586, 16458756586}
	m.caches = make([]map[string]Cache, n-4)
	m.cacheLocks = make([]sync.Mutex, n-4)
	for i := uint8(0); i < n-4; i++ {
		m.caches[i] = make(map[string]Cache)
		if isCached {
			path := fmt.Sprintf("./cache%d.gob", i+5)
			if _, err := os.Stat(path); os.IsExist(err) {
				if err := pkg.ReadGob(path, &m.caches[i]); err != nil {
					panic(err)
				}
			}
		}
		m.cacheLocks[i] = sync.Mutex{}
	}
	m.jobs = make(chan []uint8, maxJobs)
	m.workerLocks = make([]sync.Mutex, workersNumber)
	for i := uint16(0); i < workersNumber; i++ {
		m.workerLocks[i] = sync.Mutex{}
	}
	m.maxCost = (n - 2) / 3
	m.maxPerms = make([][]uint8, 0)
	m.maxLock = sync.Mutex{}
	m.countLock = sync.Mutex{}
	return &m
}

func (m *Manager) start() {
	start := time.Now()
	first := make([]uint8, 2)
	first[0] = 1
	first[1] = 2
	m.jobs <- first
	m.createWorkers()
	fmt.Println("Finished in", time.Since(start))
	m.writeResults()
}

func (m *Manager) writeResults() {
	if m.isCached {
		for i, cache := range m.caches[:m.n-5] {
			path := fmt.Sprintf("./cache%d.gob", i+5)
			newMap := make(map[string]uint8)
			for key, value := range cache {
				newMap[key] = value.cost
			}
			if err := pkg.WriteGob(path, &newMap); err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("MaxCost:", m.maxCost)
	fmt.Println("Number of canonical MaxPerms:", len(m.maxPerms))
	for _, perm := range m.maxPerms {
		fmt.Println(perm)
	}
}

// Required critical section
func (m *Manager) addSignature(perm []uint8, cost uint8, worker_id uint16) {
	eqvs := eqv.GetEqv(perm)
	cache_id := len(perm) - 5
	for _, eperm := range eqvs {
		signature := pkg.ToString(eperm)
		m.caches[cache_id][signature] = Cache{cost, worker_id}
	}
}

func (m *Manager) createWorkers() {
	for i := uint16(0); i < m.workersNumber; i++ {
		m.wgWorkers.Add(1)
		go worker(m, i)
	}
	m.wgWorkers.Wait()
}

func worker(m *Manager, worker_id uint16) {
	for perm := range m.jobs {
		// Not final level yet, so we only need to cache if non-divisible and enum
		if uint8(len(perm)) < m.n {
			// If len >= 5, may need to process for caching
			if len(perm) >= 5 {
				// It is non divisible, proceed to check if cache existed
				if div.IsMosaicDivisible(perm) == 0 {
					cacheMosaicOCS(m, worker_id, perm)
				}
			}
			output := bxt.EnumNext(perm)
			for _, newPerm := range output {
				m.jobs <- newPerm
			}
		} else { // final level
			if !sep.IsSeparable(perm) {
				cost := finalMosaicOCS(m, worker_id, perm)
				// There are possible errors in the proof of bijection between Baxter permutations and Mosaic Floorplan
				// This commented code is to check such thing
				// if eqv.IsEqv(perm, pkg.GetSlice("7 4 1 3 6 8 5 2")) {
				// 	fmt.Println(eqv.GetEqv(perm))
				// 	fmt.Println(cost, perm)
				// }
				if cost > m.maxCost {
					m.maxLock.Lock()
					if cost > m.maxCost {
						m.maxCost = cost
						m.maxPerms = make([][]uint8, 0)
						m.maxPerms = append(m.maxPerms, perm)
					} else if cost == m.maxCost {
						m.maxPerms = append(m.maxPerms, perm)
					}
					m.maxLock.Unlock()
				} else if cost == m.maxCost {
					m.maxLock.Lock()
					m.maxPerms = append(m.maxPerms, perm)
					m.maxLock.Unlock()
				}
			}
			m.countLock.Lock()
			m.count++
			if m.count == m.baxter_numbers[m.n]/2 {
				close(m.jobs)
			}
			m.countLock.Unlock()
		}
	}
	m.wgWorkers.Done()
}
