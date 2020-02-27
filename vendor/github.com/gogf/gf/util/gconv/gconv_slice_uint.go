// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

// SliceUint is alias of Uints.
func SliceUint(i interface{}) []uint {
	return Uints(i)
}

// SliceUint32 is alias of Uint32s.
func SliceUint32(i interface{}) []uint32 {
	return Uint32s(i)
}

// SliceUint64 is alias of Uint64s.
func SliceUint64(i interface{}) []uint64 {
	return Uint64s(i)
}

// Uints converts <i> to []uint.
func Uints(i interface{}) []uint {
	if i == nil {
		return nil
	}
	if r, ok := i.([]uint); ok {
		return r
	} else {
		var array []uint
		switch value := i.(type) {
		case []string:
			array = make([]uint, len(value))
			for k, v := range value {
				array[k] = Uint(v)
			}
		case []int8:
			array = make([]uint, len(value))
			for k, v := range value {
				array[k] = uint(v)
			}
		case []int16:
			array = make([]uint, len(value))
			for k, v := range value {
				array[k] = uint(v)
			}
		case []int32:
			array = make([]uint, len(value))
			for k, v := range value {
				array[k] = uint(v)
			}
		case []int64:
			array = make([]uint, len(value))
			for k, v := range value {
				array[k] = uint(v)
			}
		case []uint8:
			array = make([]uint, len(value))
			for k, v := range value {
				array[k] = uint(v)
			}
		case []uint16:
			array = make([]uint, len(value))
			for k, v := range value {
				array[k] = uint(v)
			}
		case []uint32:
			array = make([]uint, len(value))
			for k, v := range value {
				array[k] = uint(v)
			}
		case []uint64:
			array = make([]uint, len(value))
			for k, v := range value {
				array[k] = uint(v)
			}
		case []bool:
			array = make([]uint, len(value))
			for k, v := range value {
				if v {
					array[k] = 1
				} else {
					array[k] = 0
				}
			}
		case []float32:
			array = make([]uint, len(value))
			for k, v := range value {
				array[k] = Uint(v)
			}
		case []float64:
			array = make([]uint, len(value))
			for k, v := range value {
				array[k] = Uint(v)
			}
		case []interface{}:
			array = make([]uint, len(value))
			for k, v := range value {
				array[k] = Uint(v)
			}
		case [][]byte:
			array = make([]uint, len(value))
			for k, v := range value {
				array[k] = Uint(v)
			}
		default:
			return []uint{Uint(i)}
		}
		return array
	}
}

// Uint32s converts <i> to []uint32.
func Uint32s(i interface{}) []uint32 {
	if i == nil {
		return nil
	}
	if r, ok := i.([]uint32); ok {
		return r
	} else {
		var array []uint32
		switch value := i.(type) {
		case []string:
			array = make([]uint32, len(value))
			for k, v := range value {
				array[k] = Uint32(v)
			}
		case []int8:
			array = make([]uint32, len(value))
			for k, v := range value {
				array[k] = uint32(v)
			}
		case []int16:
			array = make([]uint32, len(value))
			for k, v := range value {
				array[k] = uint32(v)
			}
		case []int32:
			array = make([]uint32, len(value))
			for k, v := range value {
				array[k] = uint32(v)
			}
		case []int64:
			array = make([]uint32, len(value))
			for k, v := range value {
				array[k] = uint32(v)
			}
		case []uint:
			array = make([]uint32, len(value))
			for k, v := range value {
				array[k] = uint32(v)
			}
		case []uint8:
			array = make([]uint32, len(value))
			for k, v := range value {
				array[k] = uint32(v)
			}
		case []uint16:
			array = make([]uint32, len(value))
			for k, v := range value {
				array[k] = uint32(v)
			}
		case []uint64:
			array = make([]uint32, len(value))
			for k, v := range value {
				array[k] = uint32(v)
			}
		case []bool:
			array = make([]uint32, len(value))
			for k, v := range value {
				if v {
					array[k] = 1
				} else {
					array[k] = 0
				}
			}
		case []float32:
			array = make([]uint32, len(value))
			for k, v := range value {
				array[k] = Uint32(v)
			}
		case []float64:
			array = make([]uint32, len(value))
			for k, v := range value {
				array[k] = Uint32(v)
			}
		case []interface{}:
			array = make([]uint32, len(value))
			for k, v := range value {
				array[k] = Uint32(v)
			}
		case [][]byte:
			array = make([]uint32, len(value))
			for k, v := range value {
				array[k] = Uint32(v)
			}
		default:
			return []uint32{Uint32(i)}
		}
		return array
	}
}

// Uint64s converts <i> to []uint64.
func Uint64s(i interface{}) []uint64 {
	if i == nil {
		return nil
	}
	if r, ok := i.([]uint64); ok {
		return r
	} else {
		var array []uint64
		switch value := i.(type) {
		case []string:
			array = make([]uint64, len(value))
			for k, v := range value {
				array[k] = Uint64(v)
			}
		case []int8:
			array = make([]uint64, len(value))
			for k, v := range value {
				array[k] = uint64(v)
			}
		case []int16:
			array = make([]uint64, len(value))
			for k, v := range value {
				array[k] = uint64(v)
			}
		case []int32:
			array = make([]uint64, len(value))
			for k, v := range value {
				array[k] = uint64(v)
			}
		case []int64:
			array = make([]uint64, len(value))
			for k, v := range value {
				array[k] = uint64(v)
			}
		case []uint:
			array = make([]uint64, len(value))
			for k, v := range value {
				array[k] = uint64(v)
			}
		case []uint8:
			array = make([]uint64, len(value))
			for k, v := range value {
				array[k] = uint64(v)
			}
		case []uint16:
			array = make([]uint64, len(value))
			for k, v := range value {
				array[k] = uint64(v)
			}
		case []uint32:
			array = make([]uint64, len(value))
			for k, v := range value {
				array[k] = uint64(v)
			}
		case []bool:
			array = make([]uint64, len(value))
			for k, v := range value {
				if v {
					array[k] = 1
				} else {
					array[k] = 0
				}
			}
		case []float32:
			array = make([]uint64, len(value))
			for k, v := range value {
				array[k] = Uint64(v)
			}
		case []float64:
			array = make([]uint64, len(value))
			for k, v := range value {
				array[k] = Uint64(v)
			}
		case []interface{}:
			array = make([]uint64, len(value))
			for k, v := range value {
				array[k] = Uint64(v)
			}
		case [][]byte:
			array = make([]uint64, len(value))
			for k, v := range value {
				array[k] = Uint64(v)
			}
		default:
			return []uint64{Uint64(i)}
		}
		return array
	}
}
