package mocs

import (
	"GuillotineCuts/BP2FP"
	"GuillotineCuts/div"
	"GuillotineCuts/pkg"
	"GuillotineCuts/sep"
	"fmt"
	"sync"
	"time"
)

type Manager struct {
	caches         []map[string]uint8
	mutexes        []sync.Mutex
	n, maxCost     uint8
	maxPerm        []uint8
	workers_number uint16
	jobs           chan Job
	results        chan Result
	done           chan bool
}

func NewManager(n uint8, workers_number uint16) *Manager {
	m := Manager{n: n, workers_number: workers_number}
	m.caches = make([]map[string]uint8, n-4)
	m.mutexes = make([]sync.Mutex, n-4)
	for i := uint8(0); i < n-4; i++ {
		m.caches[i] = make(map[string]uint8)
		m.mutexes[i] = sync.Mutex{}
	}
	m.maxCost = 0
	m.jobs = make(chan Job, workers_number)
	m.results = make(chan Result, workers_number)
	m.done = make(chan bool)
	return &m
}

func (m *Manager) Start() {
	startTime := time.Now()
	go m.getResults()
	m.createWorkers()
	<-m.done
	fmt.Println("Time taken ", time.Since(startTime))
}

func (m *Manager) PrintResult() {
	fmt.Println("Perm: ", m.maxPerm)
	fmt.Println("Cost: ", m.maxCost)
}

type Job struct {
	perm []uint8
}

type Result struct {
	job  Job
	cost uint8
}

func worker(m *Manager, wg *sync.WaitGroup) {
	for job := range m.jobs {
		output := Result{job, m.mosaicOCS(job.perm)}
		m.results <- output
	}
	wg.Done()
}

func (m *Manager) createWorkers() {
	var wg sync.WaitGroup
	for i := uint16(0); i < m.workers_number; i++ {
		wg.Add(1)
		go worker(m, &wg)
	}
	wg.Wait()
	close(m.results)
}

func (m *Manager) AddJobs(perms [][]uint8) {
	for _, perm := range perms {
		job := Job{perm}
		m.jobs <- job
	}
	close(m.jobs)
}

func (m *Manager) getResults() {
	for result := range m.results {
		if m.maxCost < result.cost {
			m.maxCost = result.cost
			m.maxPerm = result.job.perm
		}
	}
	m.done <- true
}

func (m *Manager) mosaicOCS(perm []uint8) uint8 {
	// Separable mosaic
	// O(n)
	if sep.IsSeparable(perm) {
		return 0
	}

	// Divisible mosaic
	// O(nlogn)
	if cut := div.IsMosaicDivisible(perm); cut > 0 {
		return m.mosaicOCS(pkg.ToPermutation(perm[:cut])) + m.mosaicOCS(pkg.ToPermutation(perm[cut:]))
	}

	// Non-divisible mosaic
	// O(nlogn)
	// Need to cache
	signature := pkg.ToString(perm)
	m.mutexes[len(perm)-5].Lock()
	if cost, ok := m.caches[len(perm)-5][signature]; ok {
		m.mutexes[len(perm)-5].Unlock()
		return cost
	}
	m.mutexes[len(perm)-5].Unlock()
	n := uint8(len(perm))
	minCost := uint8(len(perm))
	rects := BP2FP.BP2FP(perm)
	hcuts := make(map[uint8]bool)
	vcuts := make(map[uint8]bool)
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

	for cut, _ := range vcuts {
		left := make([]uint8, 0)
		right := make([]uint8, 0)
		for _, p := range perm {
			rect := rects[p]
			if rect[0] >= cut {
				right = append(right, p)
			}
			if rect[1] <= cut {
				left = append(left, p)
			}
		}
		cost := uint8(len(perm) - len(left) - len(right))
		cost += m.mosaicOCS(pkg.ToPermutation(left)) + m.mosaicOCS(pkg.ToPermutation(right))
		minCost = pkg.Min(minCost, cost)
	}

	for cut, _ := range hcuts {
		left := make([]uint8, 0)
		right := make([]uint8, 0)
		for _, p := range perm {
			rect := rects[p]
			if rect[2] >= cut {
				right = append(right, p)
			}
			if rect[3] <= cut {
				left = append(left, p)
			}
		}
		cost := uint8(len(perm) - len(left) - len(right))
		cost += m.mosaicOCS(pkg.ToPermutation(left)) + m.mosaicOCS(pkg.ToPermutation(right))
		minCost = pkg.Min(minCost, cost)
	}

	m.mutexes[len(perm)-5].Lock()
	m.caches[len(perm)-5][signature] = minCost
	m.mutexes[len(perm)-5].Unlock()
	return minCost
}
