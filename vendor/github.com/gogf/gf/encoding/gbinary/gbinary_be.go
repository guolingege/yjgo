// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gbinary

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

// 针对基本类型进行二进制打包，支持的基本数据类型包括:
// int/8/16/32/64、uint/8/16/32/64、float32/64、bool、string、[]byte。
// 其他未知类型使用 fmt.Sprintf("%v", value) 转换为字符串之后处理。
func BeEncode(values ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for i := 0; i < len(values); i++ {
		if values[i] == nil {
			return buf.Bytes()
		}

		switch value := values[i].(type) {
		case int:
			buf.Write(BeEncodeInt(value))
		case int8:
			buf.Write(BeEncodeInt8(value))
		case int16:
			buf.Write(BeEncodeInt16(value))
		case int32:
			buf.Write(BeEncodeInt32(value))
		case int64:
			buf.Write(BeEncodeInt64(value))
		case uint:
			buf.Write(BeEncodeUint(value))
		case uint8:
			buf.Write(BeEncodeUint8(value))
		case uint16:
			buf.Write(BeEncodeUint16(value))
		case uint32:
			buf.Write(BeEncodeUint32(value))
		case uint64:
			buf.Write(BeEncodeUint64(value))
		case bool:
			buf.Write(BeEncodeBool(value))
		case string:
			buf.Write(BeEncodeString(value))
		case []byte:
			buf.Write(value)
		case float32:
			buf.Write(BeEncodeFloat32(value))
		case float64:
			buf.Write(BeEncodeFloat64(value))
		default:
			if err := binary.Write(buf, binary.BigEndian, value); err != nil {
				buf.Write(BeEncodeString(fmt.Sprintf("%v", value)))
			}
		}
	}
	return buf.Bytes()
}

// 将变量转换为二进制[]byte，并指定固定的[]byte长度返回，长度单位为字节(byte)；
// 如果转换的二进制长度超过指定长度，那么进行截断处理
func BeEncodeByLength(length int, values ...interface{}) []byte {
	b := BeEncode(values...)
	if len(b) < length {
		b = append(b, make([]byte, length-len(b))...)
	} else if len(b) > length {
		b = b[0:length]
	}
	return b
}

// 整形二进制解包，注意第二个及其后参数为字长确定的整形变量的指针地址，以便确定解析的[]byte长度，
// 例如：int8/16/32/64、uint8/16/32/64、float32/64等等
func BeDecode(b []byte, values ...interface{}) error {
	buf := bytes.NewBuffer(b)
	for i := 0; i < len(values); i++ {
		err := binary.Read(buf, binary.BigEndian, values[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func BeEncodeString(s string) []byte {
	return []byte(s)
}

func BeDecodeToString(b []byte) string {
	return string(b)
}

func BeEncodeBool(b bool) []byte {
	if b == true {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

// 自动识别int类型长度，转换为[]byte
func BeEncodeInt(i int) []byte {
	if i <= math.MaxInt8 {
		return BeEncodeInt8(int8(i))
	} else if i <= math.MaxInt16 {
		return BeEncodeInt16(int16(i))
	} else if i <= math.MaxInt32 {
		return BeEncodeInt32(int32(i))
	} else {
		return BeEncodeInt64(int64(i))
	}
}

// 自动识别uint类型长度，转换为[]byte
func BeEncodeUint(i uint) []byte {
	if i <= math.MaxUint8 {
		return BeEncodeUint8(uint8(i))
	} else if i <= math.MaxUint16 {
		return BeEncodeUint16(uint16(i))
	} else if i <= math.MaxUint32 {
		return BeEncodeUint32(uint32(i))
	} else {
		return BeEncodeUint64(uint64(i))
	}
}

func BeEncodeInt8(i int8) []byte {
	return []byte{byte(i)}
}

func BeEncodeUint8(i uint8) []byte {
	return []byte{byte(i)}
}

func BeEncodeInt16(i int16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(i))
	return b
}

func BeEncodeUint16(i uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return b
}

func BeEncodeInt32(i int32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(i))
	return b
}

func BeEncodeUint32(i uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

func BeEncodeInt64(i int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

func BeEncodeUint64(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

func BeEncodeFloat32(f float32) []byte {
	bits := math.Float32bits(f)
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, bits)
	return b
}

func BeEncodeFloat64(f float64) []byte {
	bits := math.Float64bits(f)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, bits)
	return b
}

// 将二进制解析为int类型，根据[]byte的长度进行自动转换.
// 注意内部使用的是uint*，使用int会造成位丢失。
func BeDecodeToInt(b []byte) int {
	if len(b) < 2 {
		return int(BeDecodeToUint8(b))
	} else if len(b) < 3 {
		return int(BeDecodeToUint16(b))
	} else if len(b) < 5 {
		return int(BeDecodeToUint32(b))
	} else {
		return int(BeDecodeToUint64(b))
	}
}

// 将二进制解析为uint类型，根据[]byte的长度进行自动转换
func BeDecodeToUint(b []byte) uint {
	if len(b) < 2 {
		return uint(BeDecodeToUint8(b))
	} else if len(b) < 3 {
		return uint(BeDecodeToUint16(b))
	} else if len(b) < 5 {
		return uint(BeDecodeToUint32(b))
	} else {
		return uint(BeDecodeToUint64(b))
	}
}

// 将二进制解析为bool类型，识别标准是判断二进制中数值是否都为0，或者为空。
func BeDecodeToBool(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	if bytes.Compare(b, make([]byte, len(b))) == 0 {
		return false
	}
	return true
}

func BeDecodeToInt8(b []byte) int8 {
	return int8(b[0])
}

func BeDecodeToUint8(b []byte) uint8 {
	return uint8(b[0])
}

func BeDecodeToInt16(b []byte) int16 {
	return int16(binary.BigEndian.Uint16(BeFillUpSize(b, 2)))
}

func BeDecodeToUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(BeFillUpSize(b, 2))
}

func BeDecodeToInt32(b []byte) int32 {
	return int32(binary.BigEndian.Uint32(BeFillUpSize(b, 4)))
}

func BeDecodeToUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(BeFillUpSize(b, 4))
}

func BeDecodeToInt64(b []byte) int64 {
	return int64(binary.BigEndian.Uint64(BeFillUpSize(b, 8)))
}

func BeDecodeToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(BeFillUpSize(b, 8))
}

func BeDecodeToFloat32(b []byte) float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(BeFillUpSize(b, 4)))
}

func BeDecodeToFloat64(b []byte) float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(BeFillUpSize(b, 8)))
}

// 当b位数不够时，进行低位补0。
// 注意这里为了不影响原有输入参数，是采用的值复制设计。
func BeFillUpSize(b []byte, l int) []byte {
	if len(b) >= l {
		return b[:l]
	}
	c := make([]byte, l)
	copy(c[l-len(b):], b)
	return c
}
