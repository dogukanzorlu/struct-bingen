package parser

import (
	"os"
	"unsafe"
)

// Config64 returns a default types config for 64 bit systems.
func Config64() Config {
	c := Config{PtrSize: 8, IntSize: 8}
	c.setDefaults()
	return c
}

// Config stores configuration for base types.
type Config struct {
	PtrSize     int  // size of pointers in bytes
	IntSize     int  // default int size in bytes
	WCharSize   int  // wchar_t size
	WCharSigned bool // is wchar_t signed?
	UseGoInt    bool // use Go int for C int and long
}

func (c *Config) setDefaults() {
	if c.WCharSize == 0 {
		c.WCharSize = 2
	}
	if c.PtrSize == 0 {
		switch os.Getenv("GOARCH") {
		case "386":
			c.PtrSize = 4
		case "amd64":
			c.PtrSize = 8
		default:
			c.PtrSize = int(unsafe.Sizeof((*int)(nil)))
		}
	}
	if c.IntSize == 0 {
		switch os.Getenv("GOARCH") {
		case "386":
			c.IntSize = 4
		case "amd64":
			c.IntSize = 8
		default:
			c.IntSize = int(unsafe.Sizeof(int(0)))
		}
	}
}
