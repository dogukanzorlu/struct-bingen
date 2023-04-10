package translator

import (
	"reflect"
)

var (
	uint8Type  = reflect.TypeOf(uint8(0)).Kind()
	uint16Type = reflect.TypeOf(uint16(0)).Kind()
	uint32Type = reflect.TypeOf(uint32(0)).Kind()
	uint64Type = reflect.TypeOf(uint64(0)).Kind()
	int8Type   = reflect.TypeOf(int8(0)).Kind()
	int16Type  = reflect.TypeOf(int16(0)).Kind()
	int32Type  = reflect.TypeOf(int32(0)).Kind()
	int64Type  = reflect.TypeOf(int64(0)).Kind()
)

var (
	float32Type = reflect.TypeOf(float32(0)).Kind()
	float64Type = reflect.TypeOf(float64(0)).Kind()
)

func UintT(size int) reflect.Kind {
	switch size {
	case 1:
		return uint8Type
	case 2:
		return uint16Type
	case 4:
		return uint32Type
	case 8:
		return uint64Type
	default:
		panic("undefined uint type")
	}
	return uint64Type
}

func IntT(size int) reflect.Kind {
	switch size {
	case 1:
		return int8Type
	case 2:
		return int16Type
	case 4:
		return int32Type
	case 8:
		return int64Type
	default:
		panic("undefined int type")
	}
	return int64Type
}

func FloatT(size int) reflect.Kind {
	switch size {
	case 4:
		return float32Type
	case 8:
		return float64Type
	default:
		return float64Type
	}
}
