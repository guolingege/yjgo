// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

import (
	"errors"
	"fmt"
	"reflect"
)

// SliceMap is alias of Maps.
func SliceMap(i interface{}) []map[string]interface{} {
	return Maps(i)
}

// SliceMapDeep is alias of MapsDeep.
func SliceMapDeep(i interface{}) []map[string]interface{} {
	return MapsDeep(i)
}

// SliceStruct is alias of Structs.
func SliceStruct(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return Structs(params, pointer, mapping...)
}

// SliceStructDeep is alias of StructsDeep.
func SliceStructDeep(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return StructsDeep(params, pointer, mapping...)
}

// Maps converts <i> to []map[string]interface{}.
func Maps(value interface{}, tags ...string) []map[string]interface{} {
	if value == nil {
		return nil
	}
	if r, ok := value.([]map[string]interface{}); ok {
		return r
	} else {
		array := Interfaces(value)
		if len(array) == 0 {
			return nil
		}
		list := make([]map[string]interface{}, len(array))
		for k, v := range array {
			list[k] = Map(v, tags...)
		}
		return list
	}
}

// MapsDeep converts <i> to []map[string]interface{} recursively.
func MapsDeep(value interface{}, tags ...string) []map[string]interface{} {
	if value == nil {
		return nil
	}
	if r, ok := value.([]map[string]interface{}); ok {
		return r
	} else {
		array := Interfaces(value)
		if len(array) == 0 {
			return nil
		}
		list := make([]map[string]interface{}, len(array))
		for k, v := range array {
			list[k] = MapDeep(v, tags...)
		}
		return list
	}
}

// Structs converts any slice to given struct slice.
func Structs(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return doStructs(params, pointer, false, mapping...)
}

// StructsDeep converts any slice to given struct slice recursively.
func StructsDeep(params interface{}, pointer interface{}, mapping ...map[string]string) (err error) {
	return doStructs(params, pointer, true, mapping...)
}

// doStructs converts any slice to given struct slice.
//
// The parameter <params> should be type of slice.
//
// The parameter <pointer> should be type of pointer to slice of struct.
// Note that if <pointer> is a pointer to another pointer of type of slice of struct,
// it will create the struct/pointer internally.
func doStructs(params interface{}, pointer interface{}, deep bool, mapping ...map[string]string) (err error) {
	if params == nil {
		return errors.New("params cannot be nil")
	}
	if pointer == nil {
		return errors.New("object pointer cannot be nil")
	}
	pointerRv, ok := pointer.(reflect.Value)
	if !ok {
		pointerRv = reflect.ValueOf(pointer)
		if kind := pointerRv.Kind(); kind != reflect.Ptr {
			return fmt.Errorf("pointer should be type of pointer, but got: %v", kind)
		}
	}
	rv := reflect.ValueOf(params)
	kind := rv.Kind()
	for kind == reflect.Ptr {
		rv = rv.Elem()
		kind = rv.Kind()
	}
	switch kind {
	case reflect.Slice, reflect.Array:
		// If <params> is an empty slice, no conversion.
		if rv.Len() == 0 {
			return nil
		}
		array := reflect.MakeSlice(pointerRv.Type().Elem(), rv.Len(), rv.Len())
		itemType := array.Index(0).Type()
		for i := 0; i < rv.Len(); i++ {
			if itemType.Kind() == reflect.Ptr {
				// Slice element is type pointer.
				e := reflect.New(itemType.Elem()).Elem()
				if deep {
					if err = StructDeep(rv.Index(i).Interface(), e, mapping...); err != nil {
						return err
					}
				} else {
					if err = Struct(rv.Index(i).Interface(), e, mapping...); err != nil {
						return err
					}
				}
				array.Index(i).Set(e.Addr())
			} else {
				// Slice element is not type of pointer.
				e := reflect.New(itemType).Elem()
				if deep {
					if err = StructDeep(rv.Index(i).Interface(), e, mapping...); err != nil {
						return err
					}
				} else {
					if err = Struct(rv.Index(i).Interface(), e, mapping...); err != nil {
						return err
					}
				}
				array.Index(i).Set(e)
			}
		}
		pointerRv.Elem().Set(array)
		return nil
	default:
		return fmt.Errorf("params should be type of slice, but got: %v", kind)
	}
}
