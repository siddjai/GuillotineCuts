// This enumeration is based on LTR and RTL succession rule of baxter permutation
// on page 15 of this paper https://arxiv.org/pdf/1702.04529.pdf

package enum

import (
	"GuillotineCuts/mocs"
	"GuillotineCuts/pkg"
	"fmt"
	"sync"
	"time"
)

type Manager struct {
	mutex          sync.Mutex
	n              uint8
	workers_number uint16
	jobs           chan Job
	wgWorker       sync.WaitGroup
	mmocs          *mocs.Manager
	baxter_numbers [17]uint64
	count          uint64
}

func NewManager(n uint8, workers_number uint16, maxjobs uint32) *Manager {
	m := Manager{n: n, workers_number: workers_number}
	m.count = 0
	m.baxter_numbers = [17]uint64{0, 1, 2, 6, 22, 92, 422, 2074, 10754, 58202, 326240, 1882960, 11140560, 67329992, 414499438, 2593341586, 16458756586}
	m.jobs = make(chan Job, maxjobs)
	return &m
}

func (m *Manager) StartWith(mmocs *mocs.Manager) {
	m.mmocs = mmocs
	startTime := time.Now()
	first := make([]uint8, 1)
	first[0] = 1
	m.jobs <- Job{first}
	m.createWorkers()
	fmt.Println("Time taken for Enumeration", time.Since(startTime))
	m.mmocs.Done <- true
}

type Job struct {
	perm []uint8
}

func (m *Manager) worker() {
	for job := range m.jobs {
		output := enumNext(job.perm)
		if uint8(len(job.perm)) == m.n-1 {
			go m.mmocs.AddJobs(output)
			m.mutex.Lock()
			m.count++
			if m.count == m.baxter_numbers[m.n-1] {
				go m.mmocs.StopJobs()
				close(m.jobs)
			}
			m.mutex.Unlock()
		} else {
			for _, perm := range output {
				m.jobs <- Job{perm}
			}
		}
	}
	m.wgWorker.Done()
}

func (m *Manager) createWorkers() {
	for i := uint16(0); i < m.workers_number; i++ {
		m.wgWorker.Add(1)
		go m.worker()
	}
	m.wgWorker.Wait()
}

func enumNext(perm []uint8) [][]uint8 {
	n := uint8(len(perm))
	perms := make([][]uint8, 0)
	max := uint8(0)
	for i, p := range perm {
		if p > max {
			max = p
			newperm := pkg.InsertPerm(perm, n+1, uint8(i))
			perms = append(perms, newperm)
		}
	}
	max = uint8(0)
	for r := uint8(0); r < n; r++ {
		i := n - r - 1
		p := perm[i]
		if p > max {
			max = p
			newperm := pkg.InsertPerm(perm, n+1, uint8(i+1))
			perms = append(perms, newperm)
		}
	}
	return perms
}
