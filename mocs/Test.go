package mocs

func Test(n uint8, workersNumber uint16, maxJobs uint32, isCached bool) {
	m := NewManager(n, workersNumber, maxJobs, isCached)
	m.start()
}
