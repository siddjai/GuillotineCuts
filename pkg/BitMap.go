// Bitmap just in case we will do the general case

package pkg

import (
	"strconv"
)

type Bitmap struct {
	data []uint8
	size uint8 // bitsize
}

func NewBitmap(size uint8) *Bitmap {
	if remainder := size % 8; remainder != 0 {
		size += 8 - remainder
	}
	return &Bitmap{make([]uint8, size>>3), size}
}

func (this *Bitmap) SetBit(offset uint8, bit uint8) bool {
	index, pos := offset>>3, offset%8

	if this.size < offset {
		return false
	}

	if bit == 0 {
		this.data[index] &^= 0x01 << pos
	} else {
		this.data[index] |= 0x01 << pos
	}
	return true
}

func (this *Bitmap) GetBit(offset uint8) uint8 {
	index, pos := offset/8, offset%8
	if this.size < offset {
		return 0
	}
	return (this.data[index] >> pos) & 0x01
}

func (this *Bitmap) String() string {
	s := ""
	for i := this.size - 1; i > 0; i-- {
		s += strconv.Itoa(int(this.GetBit(i)))
	}
	s += strconv.Itoa(int(this.GetBit(0)))
	return s
}
