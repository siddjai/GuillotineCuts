package enum

import (
	"GuillotineCuts/mocs"
)

func Test(n uint8, workers_enum uint16, workers_mocs uint16, maxjobs uint32) {
	menum := NewManager(n, workers_enum, maxjobs)
	mmocs := mocs.NewManager(n, workers_mocs, maxjobs)
	go menum.StartWith(mmocs)
	mmocs.Start()
	mmocs.PrintResult()
}
