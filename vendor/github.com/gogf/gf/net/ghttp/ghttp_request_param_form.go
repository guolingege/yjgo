// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import (
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/internal/structs"
	"github.com/gogf/gf/util/gconv"
)

// SetForm sets custom form value with key-value pair.
func (r *Request) SetForm(key string, value interface{}) {
	r.parseForm()
	if r.formMap == nil {
		r.formMap = make(map[string]interface{})
	}
	r.formMap[key] = value
}

// GetForm retrieves and returns parameter <key> from form.
// It returns <def> if <key> does not exist in the form.
// It returns nil if <def> is not passed.
func (r *Request) GetForm(key string, def ...interface{}) interface{} {
	r.parseForm()
	if len(r.formMap) > 0 {
		if v, ok := r.formMap[key]; ok {
			return v
		}
	}
	if len(def) > 0 {
		return def[0]
	}
	return nil
}

func (r *Request) GetFormVar(key string, def ...interface{}) *gvar.Var {
	return gvar.New(r.GetForm(key, def...))
}

func (r *Request) GetFormString(key string, def ...interface{}) string {
	return r.GetFormVar(key, def...).String()
}

func (r *Request) GetFormBool(key string, def ...interface{}) bool {
	return r.GetFormVar(key, def...).Bool()
}

func (r *Request) GetFormInt(key string, def ...interface{}) int {
	return r.GetFormVar(key, def...).Int()
}

func (r *Request) GetFormInt32(key string, def ...interface{}) int32 {
	return r.GetFormVar(key, def...).Int32()
}

func (r *Request) GetFormInt64(key string, def ...interface{}) int64 {
	return r.GetFormVar(key, def...).Int64()
}

func (r *Request) GetFormInts(key string, def ...interface{}) []int {
	return r.GetFormVar(key, def...).Ints()
}

func (r *Request) GetFormUint(key string, def ...interface{}) uint {
	return r.GetFormVar(key, def...).Uint()
}

func (r *Request) GetFormUint32(key string, def ...interface{}) uint32 {
	return r.GetFormVar(key, def...).Uint32()
}

func (r *Request) GetFormUint64(key string, def ...interface{}) uint64 {
	return r.GetFormVar(key, def...).Uint64()
}

func (r *Request) GetFormFloat32(key string, def ...interface{}) float32 {
	return r.GetFormVar(key, def...).Float32()
}

func (r *Request) GetFormFloat64(key string, def ...interface{}) float64 {
	return r.GetFormVar(key, def...).Float64()
}

func (r *Request) GetFormFloats(key string, def ...interface{}) []float64 {
	return r.GetFormVar(key, def...).Floats()
}

func (r *Request) GetFormArray(key string, def ...interface{}) []string {
	return r.GetFormVar(key, def...).Strings()
}

func (r *Request) GetFormStrings(key string, def ...interface{}) []string {
	return r.GetFormVar(key, def...).Strings()
}

func (r *Request) GetFormInterfaces(key string, def ...interface{}) []interface{} {
	return r.GetFormVar(key, def...).Interfaces()
}

// GetFormMap retrieves and returns all form parameters passed from client as map.
// The parameter <kvMap> specifies the keys retrieving from client parameters,
// the associated values are the default values if the client does not pass.
func (r *Request) GetFormMap(kvMap ...map[string]interface{}) map[string]interface{} {
	r.parseForm()
	if len(kvMap) > 0 && kvMap[0] != nil {
		if len(r.formMap) == 0 {
			return kvMap[0]
		}
		m := make(map[string]interface{}, len(kvMap[0]))
		for k, defValue := range kvMap[0] {
			if postValue, ok := r.formMap[k]; ok {
				m[k] = postValue
			} else {
				m[k] = defValue
			}
		}
		return m
	} else {
		return r.formMap
	}
}

// GetFormMapStrStr retrieves and returns all form parameters passed from client as map[string]string.
// The parameter <kvMap> specifies the keys retrieving from client parameters, the associated values
// are the default values if the client does not pass.
func (r *Request) GetFormMapStrStr(kvMap ...map[string]interface{}) map[string]string {
	postMap := r.GetFormMap(kvMap...)
	if len(postMap) > 0 {
		m := make(map[string]string, len(postMap))
		for k, v := range postMap {
			m[k] = gconv.String(v)
		}
		return m
	}
	return nil
}

// GetFormMapStrVar retrieves and returns all form parameters passed from client as map[string]*gvar.Var.
// The parameter <kvMap> specifies the keys retrieving from client parameters, the associated values
// are the default values if the client does not pass.
func (r *Request) GetFormMapStrVar(kvMap ...map[string]interface{}) map[string]*gvar.Var {
	postMap := r.GetFormMap(kvMap...)
	if len(postMap) > 0 {
		m := make(map[string]*gvar.Var, len(postMap))
		for k, v := range postMap {
			m[k] = gvar.New(v)
		}
		return m
	}
	return nil
}

// GetFormStruct retrieves all form parameters passed from client and converts them to
// given struct object. Note that the parameter <pointer> is a pointer to the struct object.
// The optional parameter <mapping> is used to specify the key to attribute mapping.
func (r *Request) GetFormStruct(pointer interface{}, mapping ...map[string]string) error {
	r.parseForm()
	tagMap := structs.TagMapName(pointer, paramTagPriority, true)
	if len(mapping) > 0 {
		for k, v := range mapping[0] {
			tagMap[k] = v
		}
	}
	return gconv.StructDeep(r.formMap, pointer, tagMap)
}

// GetFormToStruct is alias of GetFormStruct. See GetFormStruct.
// Deprecated.
func (r *Request) GetFormToStruct(pointer interface{}, mapping ...map[string]string) error {
	return r.GetFormStruct(pointer, mapping...)
}
