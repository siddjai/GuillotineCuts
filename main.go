package main

func main() {
	n := 11
	workersNumber := 1000
	maxJobs := 10000000
	mocs.Test(uint8(n), uint16(workersNumber), uint32(maxJobs), false)
}