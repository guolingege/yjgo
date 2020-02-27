// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gjson provides convenient API for JSON/XML/YAML/TOML data handling.
package gjson

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/gogf/gf/internal/rwmutex"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

const (
	// Separator char for hierarchical data access.
	gDEFAULT_SPLIT_CHAR = '.'
)

// The customized JSON struct.
type Json struct {
	mu *rwmutex.RWMutex
	p  *interface{} // Pointer for hierarchical data access, it's the root of data in default.
	c  byte         // Char separator('.' in default).
	vc bool         // Violence Check(false in default), which is used to access data when the hierarchical data key contains separator char.
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (j *Json) MarshalJSON() ([]byte, error) {
	return j.ToJson()
}

// setValue sets <value> to <j> by <pattern>.
// Note:
// 1. If value is nil and removed is true, means deleting this value;
// 2. It's quite complicated in hierarchical data search, node creating and data assignment;
func (j *Json) setValue(pattern string, value interface{}, removed bool) error {
	array := strings.Split(pattern, string(j.c))
	length := len(array)
	value = j.convertValue(value)
	// 初始化判断
	if *j.p == nil {
		if gstr.IsNumeric(array[0]) {
			*j.p = make([]interface{}, 0)
		} else {
			*j.p = make(map[string]interface{})
		}
	}
	var pparent *interface{} = nil // Parent pointer.
	var pointer *interface{} = j.p // Current pointer.
	j.mu.Lock()
	defer j.mu.Unlock()
	for i := 0; i < length; i++ {
		switch (*pointer).(type) {
		case map[string]interface{}:
			if i == length-1 {
				if removed && value == nil {
					// Delete item from map.
					delete((*pointer).(map[string]interface{}), array[i])
				} else {
					(*pointer).(map[string]interface{})[array[i]] = value
				}
			} else {
				// If the key does not exit in the map.
				if v, ok := (*pointer).(map[string]interface{})[array[i]]; !ok {
					if removed && value == nil {
						goto done
					}
					// Creating new node.
					if gstr.IsNumeric(array[i+1]) {
						// Creating array node.
						n, _ := strconv.Atoi(array[i+1])
						var v interface{} = make([]interface{}, n+1)
						pparent = j.setPointerWithValue(pointer, array[i], v)
						pointer = &v
					} else {
						// Creating map node.
						var v interface{} = make(map[string]interface{})
						pparent = j.setPointerWithValue(pointer, array[i], v)
						pointer = &v
					}
				} else {
					pparent = pointer
					pointer = &v
				}
			}

		case []interface{}:
			if !gstr.IsNumeric(array[i]) {
				if i == length-1 {
					*pointer = map[string]interface{}{array[i]: value}
				} else {
					var v interface{} = make(map[string]interface{})
					*pointer = v
					pparent = pointer
					pointer = &v
				}
				continue
			}

			valn, err := strconv.Atoi(array[i])
			if err != nil {
				return err
			}
			// Leaf node.
			if i == length-1 {
				if len((*pointer).([]interface{})) > valn {
					if removed && value == nil {
						// Deleting element.
						if pparent == nil {
							*pointer = append((*pointer).([]interface{})[:valn], (*pointer).([]interface{})[valn+1:]...)
						} else {
							j.setPointerWithValue(pparent, array[i-1], append((*pointer).([]interface{})[:valn], (*pointer).([]interface{})[valn+1:]...))
						}
					} else {
						(*pointer).([]interface{})[valn] = value
					}
				} else {
					if removed && value == nil {
						goto done
					}
					if pparent == nil {
						// It is the root node.
						j.setPointerWithValue(pointer, array[i], value)
					} else {
						// It is not the root node.
						s := make([]interface{}, valn+1)
						copy(s, (*pointer).([]interface{}))
						s[valn] = value
						j.setPointerWithValue(pparent, array[i-1], s)
					}
				}
			} else {
				if gstr.IsNumeric(array[i+1]) {
					n, _ := strconv.Atoi(array[i+1])
					if len((*pointer).([]interface{})) > valn {
						(*pointer).([]interface{})[valn] = make([]interface{}, n+1)
						pparent = pointer
						pointer = &(*pointer).([]interface{})[valn]
					} else {
						if removed && value == nil {
							goto done
						}
						var v interface{} = make([]interface{}, n+1)
						pparent = j.setPointerWithValue(pointer, array[i], v)
						pointer = &v
					}
				} else {
					var v interface{} = make(map[string]interface{})
					pparent = j.setPointerWithValue(pointer, array[i], v)
					pointer = &v
				}
			}

		// If the variable pointed to by the <pointer> is not of a reference type,
		// then it modifies the variable via its the parent, ie: pparent.
		default:
			if removed && value == nil {
				goto done
			}
			if gstr.IsNumeric(array[i]) {
				n, _ := strconv.Atoi(array[i])
				s := make([]interface{}, n+1)
				if i == length-1 {
					s[n] = value
				}
				if pparent != nil {
					pparent = j.setPointerWithValue(pparent, array[i-1], s)
				} else {
					*pointer = s
					pparent = pointer
				}
			} else {
				var v interface{} = make(map[string]interface{})
				if i == length-1 {
					v = map[string]interface{}{
						array[i]: value,
					}
				}
				if pparent != nil {
					pparent = j.setPointerWithValue(pparent, array[i-1], v)
				} else {
					*pointer = v
					pparent = pointer
				}
				pointer = &v
			}
		}
	}
done:
	return nil
}

// convertValue converts <value> to map[string]interface{} or []interface{},
// which can be supported for hierarchical data access.
func (j *Json) convertValue(value interface{}) interface{} {
	switch value.(type) {
	case map[string]interface{}:
		return value
	case []interface{}:
		return value
	default:
		rv := reflect.ValueOf(value)
		kind := rv.Kind()
		if kind == reflect.Ptr {
			rv = rv.Elem()
			kind = rv.Kind()
		}
		switch kind {
		case reflect.Array:
			return gconv.Interfaces(value)
		case reflect.Slice:
			return gconv.Interfaces(value)
		case reflect.Map:
			return gconv.Map(value)
		case reflect.Struct:
			return gconv.Map(value)
		default:
			// Use json decode/encode at last.
			b, _ := Encode(value)
			v, _ := Decode(b)
			return v
		}
	}
}

// setPointerWithValue sets <key>:<value> to <pointer>, the <key> may be a map key or slice index.
// It returns the pointer to the new value set.
func (j *Json) setPointerWithValue(pointer *interface{}, key string, value interface{}) *interface{} {
	switch (*pointer).(type) {
	case map[string]interface{}:
		(*pointer).(map[string]interface{})[key] = value
		return &value
	case []interface{}:
		n, _ := strconv.Atoi(key)
		if len((*pointer).([]interface{})) > n {
			(*pointer).([]interface{})[n] = value
			return &(*pointer).([]interface{})[n]
		} else {
			s := make([]interface{}, n+1)
			copy(s, (*pointer).([]interface{}))
			s[n] = value
			*pointer = s
			return &s[n]
		}
	default:
		*pointer = value
	}
	return pointer
}

// getPointerByPattern returns a pointer to the value by specified <pattern>.
func (j *Json) getPointerByPattern(pattern string) *interface{} {
	if j.vc {
		return j.getPointerByPatternWithViolenceCheck(pattern)
	} else {
		return j.getPointerByPatternWithoutViolenceCheck(pattern)
	}
}

// getPointerByPatternWithViolenceCheck returns a pointer to the value of specified <pattern> with violence check.
func (j *Json) getPointerByPatternWithViolenceCheck(pattern string) *interface{} {
	if !j.vc {
		return j.getPointerByPatternWithoutViolenceCheck(pattern)
	}
	index := len(pattern)
	start := 0
	length := 0
	pointer := j.p
	if index == 0 {
		return pointer
	}
	for {
		if r := j.checkPatternByPointer(pattern[start:index], pointer); r != nil {
			length += index - start
			if start > 0 {
				length += 1
			}
			start = index + 1
			index = len(pattern)
			if length == len(pattern) {
				return r
			} else {
				pointer = r
			}
		} else {
			// Get the position for next separator char.
			index = strings.LastIndexByte(pattern[start:index], j.c)
			if index != -1 && length > 0 {
				index += length + 1
			}
		}
		if start >= index {
			break
		}
	}
	return nil
}

// getPointerByPatternWithoutViolenceCheck returns a pointer to the value of specified <pattern>, with no violence check.
func (j *Json) getPointerByPatternWithoutViolenceCheck(pattern string) *interface{} {
	if j.vc {
		return j.getPointerByPatternWithViolenceCheck(pattern)
	}
	pointer := j.p
	if len(pattern) == 0 {
		return pointer
	}
	array := strings.Split(pattern, string(j.c))
	for k, v := range array {
		if r := j.checkPatternByPointer(v, pointer); r != nil {
			if k == len(array)-1 {
				return r
			} else {
				pointer = r
			}
		} else {
			break
		}
	}
	return nil
}

// checkPatternByPointer checks whether there's value by <key> in specified <pointer>.
// It returns a pointer to the value.
func (j *Json) checkPatternByPointer(key string, pointer *interface{}) *interface{} {
	switch (*pointer).(type) {
	case map[string]interface{}:
		if v, ok := (*pointer).(map[string]interface{})[key]; ok {
			return &v
		}
	case []interface{}:
		if gstr.IsNumeric(key) {
			n, err := strconv.Atoi(key)
			if err == nil && len((*pointer).([]interface{})) > n {
				return &(*pointer).([]interface{})[n]
			}
		}
	}
	return nil
}
