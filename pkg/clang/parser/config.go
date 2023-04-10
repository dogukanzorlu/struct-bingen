package parser

import (
	"encoding/binary"
	"runtime"
	"unsafe"
)

type Config struct {
	PtrSize    int
	IntSize    int
	EndianType binary.ByteOrder
}

func Config64() Config {
	c := Config{}
	c.setSize()
	c.setEndian()
	return c
}

func (c *Config) setSize() {
	switch runtime.GOARCH {
	case "386":
		c.PtrSize = 4
		c.IntSize = 4
	case "arm64":
		c.PtrSize = 4
		c.IntSize = 4
	case "amd64":
		c.PtrSize = 8
		c.IntSize = 8
	default:
		c.PtrSize = int(unsafe.Sizeof((*int)(nil)))
		c.IntSize = int(unsafe.Sizeof(int(0)))
	}
}

func (c *Config) setEndian() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		c.EndianType = binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		c.EndianType = binary.BigEndian
	default:
		panic("Could not determine native endianness.")
	}
}
