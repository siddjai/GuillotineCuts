package main

import "GuillotineCuts/enum"

func main() {
	n := uint8(12)
	workers_enum := uint16(200)
	workers_mocs := uint16(500)
	maxjobs := uint32(10000000)
	enum.Test(n, workers_enum, workers_mocs, maxjobs)
}
