// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gconv

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gogf/gf/internal/empty"
	"github.com/gogf/gf/internal/utilstr"
)

// apiMapStrAny is the interface support for package gmap.
type apiMapStrAny interface {
	MapStrAny() map[string]interface{}
}

// Map converts any variable <value> to map[string]interface{}. If the parameter <value> is not a
// map/struct/*struct type, then the conversion will fail and returns nil.
//
// If <value> is a struct/*struct object, the second parameter <tags> specifies the most priority
// tags that will be detected, otherwise it detects the tags in order of: gconv, json, and then the field name.
func Map(value interface{}, tags ...string) map[string]interface{} {
	return doMapConvert(value, false, tags...)
}

// MapDeep does Map function recursively, which means if the attribute of <value>
// is also a struct/*struct, calls Map function on this attribute converting it to
// a map[string]interface{} type variable.
// Also see Map.
func MapDeep(value interface{}, tags ...string) map[string]interface{} {
	return doMapConvert(value, true, tags...)
}

// doMapConvert implements the map converting.
func doMapConvert(value interface{}, recursive bool, tags ...string) map[string]interface{} {
	if value == nil {
		return nil
	}
	if r, ok := value.(map[string]interface{}); ok {
		return r
	} else {
		// Assert the common combination of types, and finally it uses reflection.
		m := make(map[string]interface{})
		switch value.(type) {
		case map[interface{}]interface{}:
			for k, v := range value.(map[interface{}]interface{}) {
				m[String(k)] = v
			}
		case map[interface{}]string:
			for k, v := range value.(map[interface{}]string) {
				m[String(k)] = v
			}
		case map[interface{}]int:
			for k, v := range value.(map[interface{}]int) {
				m[String(k)] = v
			}
		case map[interface{}]uint:
			for k, v := range value.(map[interface{}]uint) {
				m[String(k)] = v
			}
		case map[interface{}]float32:
			for k, v := range value.(map[interface{}]float32) {
				m[String(k)] = v
			}
		case map[interface{}]float64:
			for k, v := range value.(map[interface{}]float64) {
				m[String(k)] = v
			}
		case map[string]bool:
			for k, v := range value.(map[string]bool) {
				m[k] = v
			}
		case map[string]int:
			for k, v := range value.(map[string]int) {
				m[k] = v
			}
		case map[string]uint:
			for k, v := range value.(map[string]uint) {
				m[k] = v
			}
		case map[string]float32:
			for k, v := range value.(map[string]float32) {
				m[k] = v
			}
		case map[string]float64:
			for k, v := range value.(map[string]float64) {
				m[k] = v
			}
		case map[int]interface{}:
			for k, v := range value.(map[int]interface{}) {
				m[String(k)] = v
			}
		case map[int]string:
			for k, v := range value.(map[int]string) {
				m[String(k)] = v
			}
		case map[uint]string:
			for k, v := range value.(map[uint]string) {
				m[String(k)] = v
			}
		// Not a common type, then use reflection.
		default:
			rv := reflect.ValueOf(value)
			kind := rv.Kind()
			// If it is a pointer, we should find its real data type.
			if kind == reflect.Ptr {
				rv = rv.Elem()
				kind = rv.Kind()
			}
			switch kind {
			case reflect.Map:
				ks := rv.MapKeys()
				for _, k := range ks {
					m[String(k.Interface())] = rv.MapIndex(k).Interface()
				}
			case reflect.Struct:
				if v, ok := value.(apiMapStrAny); ok {
					return v.MapStrAny()
				}
				rt := rv.Type()
				name := ""
				tagArray := structTagPriority
				switch len(tags) {
				case 0:
					// No need handle.
				case 1:
					tagArray = append(strings.Split(tags[0], ","), structTagPriority...)
				default:
					tagArray = append(tags, structTagPriority...)
				}
				var rtField reflect.StructField
				var rvField reflect.Value
				var rvKind reflect.Kind
				for i := 0; i < rv.NumField(); i++ {
					rtField = rt.Field(i)
					rvField = rv.Field(i)
					// Only convert the public attributes.
					fieldName := rtField.Name
					if !utilstr.IsLetterUpper(fieldName[0]) {
						continue
					}
					name = ""
					fieldTag := rtField.Tag
					for _, tag := range tagArray {
						if name = fieldTag.Get(tag); name != "" {
							break
						}
					}
					if name == "" {
						name = strings.TrimSpace(fieldName)
					} else {
						// Support json tag feature: -, omitempty
						name = strings.TrimSpace(name)
						if name == "-" {
							continue
						}
						array := strings.Split(name, ",")
						if len(array) > 1 {
							switch strings.TrimSpace(array[1]) {
							case "omitempty":
								if empty.IsEmpty(rvField.Interface()) {
									continue
								} else {
									name = strings.TrimSpace(array[0])
								}
							default:
								name = strings.TrimSpace(array[0])
							}
						}
					}
					if recursive {
						rvKind = rvField.Kind()
						if rvKind == reflect.Ptr {
							rvField = rvField.Elem()
							rvKind = rvField.Kind()
						}
						if rvKind == reflect.Struct {
							for k, v := range doMapConvert(rvField.Interface(), recursive, tags...) {
								m[k] = v
							}
						} else {
							m[name] = rvField.Interface()
						}
					} else {
						m[name] = rvField.Interface()
					}
				}
			default:
				return nil
			}
		}
		return m
	}
}

// MapStrStr converts <value> to map[string]string.
// Note that there might be data copy for this map type converting.
func MapStrStr(value interface{}, tags ...string) map[string]string {
	if r, ok := value.(map[string]string); ok {
		return r
	}
	m := Map(value, tags...)
	if len(m) > 0 {
		vMap := make(map[string]string)
		for k, v := range m {
			vMap[k] = String(v)
		}
		return vMap
	}
	return nil
}

// MapStrStrDeep converts <value> to map[string]string recursively.
// Note that there might be data copy for this map type converting.
func MapStrStrDeep(value interface{}, tags ...string) map[string]string {
	if r, ok := value.(map[string]string); ok {
		return r
	}
	m := MapDeep(value, tags...)
	if len(m) > 0 {
		vMap := make(map[string]string)
		for k, v := range m {
			vMap[k] = String(v)
		}
		return vMap
	}
	return nil
}

// MapToMap converts map type variable <params> to another map type variable <pointer> using reflect.
// The elements of <pointer> should be type of *map.
func MapToMap(params interface{}, pointer interface{}, mapping ...map[string]string) error {
	return doMapToMap(params, pointer, false, mapping...)
}

// MapToMapDeep recursively converts map type variable <params> to another map type variable <pointer> using reflect.
// The elements of <pointer> should be type of *map.
func MapToMapDeep(params interface{}, pointer interface{}, mapping ...map[string]string) error {
	return doMapToMap(params, pointer, true, mapping...)
}

// doMapToMap converts map type variable <params> to another map type variable <pointer>.
// The elements of <pointer> should be type of *map.
func doMapToMap(params interface{}, pointer interface{}, deep bool, mapping ...map[string]string) error {
	paramsRv := reflect.ValueOf(params)
	paramsKind := paramsRv.Kind()
	if paramsKind == reflect.Ptr {
		paramsRv = paramsRv.Elem()
		paramsKind = paramsRv.Kind()
	}
	if paramsKind != reflect.Map {
		return errors.New("params should be type of map")
	}

	pointerRv := reflect.ValueOf(pointer)
	pointerKind := pointerRv.Kind()
	for pointerKind == reflect.Ptr {
		pointerRv = pointerRv.Elem()
		pointerKind = pointerRv.Kind()
	}
	if pointerKind != reflect.Map {
		return errors.New("pointer should be type of map")
	}
	err := (error)(nil)
	paramsKeys := paramsRv.MapKeys()
	pointerKeyType := pointerRv.Type().Key()
	pointerValueType := pointerRv.Type().Elem()
	dataMap := reflect.MakeMapWithSize(pointerRv.Type(), len(paramsKeys))
	for _, key := range paramsKeys {
		e := reflect.New(pointerValueType).Elem()
		if deep {
			if err = StructDeep(paramsRv.MapIndex(key).Interface(), e, mapping...); err != nil {
				return err
			}
		} else {
			if err = Struct(paramsRv.MapIndex(key).Interface(), e, mapping...); err != nil {
				return err
			}
		}
		dataMap.SetMapIndex(
			reflect.ValueOf(Convert(key.Interface(), pointerKeyType.Name())),
			e,
		)
	}
	pointerRv.Set(dataMap)
	return nil
}

// MapToMaps converts map type variable <params> to another map type variable <pointer>.
// The elements of <pointer> should be type of []map/*map.
func MapToMaps(params interface{}, pointer interface{}, mapping ...map[string]string) error {
	return doMapToMaps(params, pointer, false, mapping...)
}

// MapToMapsDeep recursively converts map type variable <params> to another map type variable <pointer>.
// The elements of <pointer> should be type of []map/*map.
func MapToMapsDeep(params interface{}, pointer interface{}, mapping ...map[string]string) error {
	return doMapToMaps(params, pointer, true, mapping...)
}

// doMapToMaps converts map type variable <params> to another map type variable <pointer>.
// The elements of <pointer> should be type of []map/*map.
func doMapToMaps(params interface{}, pointer interface{}, deep bool, mapping ...map[string]string) error {
	paramsRv := reflect.ValueOf(params)
	paramsKind := paramsRv.Kind()
	if paramsKind == reflect.Ptr {
		paramsRv = paramsRv.Elem()
		paramsKind = paramsRv.Kind()
	}
	if paramsKind != reflect.Map {
		return errors.New("params should be type of map")
	}

	pointerRv := reflect.ValueOf(pointer)
	pointerKind := pointerRv.Kind()
	for pointerKind == reflect.Ptr {
		pointerRv = pointerRv.Elem()
		pointerKind = pointerRv.Kind()
	}
	if pointerKind != reflect.Map {
		return errors.New("pointer should be type of map")
	}
	err := (error)(nil)
	paramsKeys := paramsRv.MapKeys()
	pointerKeyType := pointerRv.Type().Key()
	pointerValueType := pointerRv.Type().Elem()
	dataMap := reflect.MakeMapWithSize(pointerRv.Type(), len(paramsKeys))
	for _, key := range paramsKeys {
		e := reflect.New(pointerValueType).Elem().Addr()
		if deep {
			if err = StructsDeep(paramsRv.MapIndex(key).Interface(), e, mapping...); err != nil {
				return err
			}
		} else {
			if err = Structs(paramsRv.MapIndex(key).Interface(), e, mapping...); err != nil {
				return err
			}
		}
		dataMap.SetMapIndex(
			reflect.ValueOf(Convert(key.Interface(), pointerKeyType.Name())),
			e.Elem(),
		)
	}
	pointerRv.Set(dataMap)
	return nil
}
