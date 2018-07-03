package pkg

import (
	"fmt"
	"sync"
	"time"
)

type Manager struct {
	caches         []map[string]uint8
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
}

func (m *Manager) StopJobs() {
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
	if IsSeparable(perm) {
		return 0
	}

	// Divisible mosaic
	// O(nlogn)
	if cut := IsMosaicDivisible(perm); cut > 0 {
		return m.mosaicOCS(ToPermutation(perm[:cut])) + m.mosaicOCS(ToPermutation(perm[cut:]))
	}

	// Non-divisible mosaic
	// O(n^2logn) Costly!!
	// Need to cache
	signature := ToString(perm)
	if cost, ok := m.caches[len(perm)-5][signature]; ok {
		return cost
	}
	cuts := BP2Cuts(perm)
	minCost := uint8(len(perm))
	for i := 1; i < len(cuts); i += 2 {
		left := make([]uint8, 0)
		right := make([]uint8, 0)
		if cuts[i-1] < cuts[i+1] {
			for j := 0; j < i; j += 2 {
				if cuts[j] < cuts[i] {
					left = append(left, cuts[j])
				}
			}
			for j := i + 1; j < len(cuts); j += 2 {
				if cuts[j] > cuts[i] {
					right = append(right, cuts[j])
				}
			}
		} else {
			for j := 0; j < i; j += 2 {
				if cuts[j] > cuts[i] {
					left = append(left, cuts[j])
				}
			}
			for j := i + 1; j < len(cuts); j += 2 {
				if cuts[j] < cuts[i] {
					right = append(right, cuts[j])
				}
			}
		}
		cost := uint8(len(perm) - len(left) - len(right))
		cost += m.mosaicOCS(ToPermutation(left)) + m.mosaicOCS(ToPermutation(right))
		minCost = Min(minCost, cost)
	}
	m.caches[len(perm)-5][signature] = minCost
	return minCost
}
