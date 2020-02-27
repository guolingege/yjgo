// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.
//

package gset

import (
	"bytes"
	"encoding/json"
	"github.com/gogf/gf/internal/rwmutex"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

type StrSet struct {
	mu   *rwmutex.RWMutex
	data map[string]struct{}
}

// New create and returns a new set, which contains un-repeated items.
// The parameter <safe> is used to specify whether using set in concurrent-safety,
// which is false in default.
func NewStrSet(safe ...bool) *StrSet {
	return &StrSet{
		mu:   rwmutex.New(safe...),
		data: make(map[string]struct{}),
	}
}

// NewStrSetFrom returns a new set from <items>.
func NewStrSetFrom(items []string, safe ...bool) *StrSet {
	m := make(map[string]struct{})
	for _, v := range items {
		m[v] = struct{}{}
	}
	return &StrSet{
		mu:   rwmutex.New(safe...),
		data: m,
	}
}

// Iterator iterates the set with given callback function <f>,
// if <f> returns true then continue iterating; or false to stop.
func (set *StrSet) Iterator(f func(v string) bool) *StrSet {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.data {
		if !f(k) {
			break
		}
	}
	return set
}

// Add adds one or multiple items to the set.
func (set *StrSet) Add(item ...string) *StrSet {
	set.mu.Lock()
	for _, v := range item {
		set.data[v] = struct{}{}
	}
	set.mu.Unlock()
	return set
}

// AddIfNotExistFunc adds the returned value of callback function <f> to the set
// if <item> does not exit in the set.
func (set *StrSet) AddIfNotExistFunc(item string, f func() string) *StrSet {
	if !set.Contains(item) {
		set.doAddWithLockCheck(item, f())
	}
	return set
}

// AddIfNotExistFuncLock adds the returned value of callback function <f> to the set
// if <item> does not exit in the set.
//
// Note that the callback function <f> is executed in the mutex.Lock of the set.
func (set *StrSet) AddIfNotExistFuncLock(item string, f func() string) *StrSet {
	if !set.Contains(item) {
		set.doAddWithLockCheck(item, f)
	}
	return set
}

// doAddWithLockCheck checks whether item exists with mutex.Lock,
// if not exists, it adds item to the set or else just returns the existing value.
//
// If <value> is type of <func() interface {}>,
// it will be executed with mutex.Lock of the set,
// and its return value will be added to the set.
//
// It returns item successfully added..
func (set *StrSet) doAddWithLockCheck(item string, value interface{}) string {
	set.mu.Lock()
	defer set.mu.Unlock()
	if _, ok := set.data[item]; !ok && value != nil {
		if f, ok := value.(func() string); ok {
			item = f()
		} else {
			item = value.(string)
		}
	}
	if item != "" {
		set.data[item] = struct{}{}
	}
	return item
}

// Contains checks whether the set contains <item>.
func (set *StrSet) Contains(item string) bool {
	set.mu.RLock()
	_, exists := set.data[item]
	set.mu.RUnlock()
	return exists
}

// Remove deletes <item> from set.
func (set *StrSet) Remove(item string) *StrSet {
	set.mu.Lock()
	delete(set.data, item)
	set.mu.Unlock()
	return set
}

// Size returns the size of the set.
func (set *StrSet) Size() int {
	set.mu.RLock()
	l := len(set.data)
	set.mu.RUnlock()
	return l
}

// Clear deletes all items of the set.
func (set *StrSet) Clear() *StrSet {
	set.mu.Lock()
	set.data = make(map[string]struct{})
	set.mu.Unlock()
	return set
}

// Slice returns the a of items of the set as slice.
func (set *StrSet) Slice() []string {
	set.mu.RLock()
	ret := make([]string, len(set.data))
	i := 0
	for item := range set.data {
		ret[i] = item
		i++
	}

	set.mu.RUnlock()
	return ret
}

// Join joins items with a string <glue>.
func (set *StrSet) Join(glue string) string {
	set.mu.RLock()
	defer set.mu.RUnlock()
	buffer := bytes.NewBuffer(nil)
	l := len(set.data)
	i := 0
	for k, _ := range set.data {
		buffer.WriteString(k)
		if i != l-1 {
			buffer.WriteString(glue)
		}
		i++
	}
	return buffer.String()
}

// String returns items as a string, which implements like json.Marshal does.
func (set *StrSet) String() string {
	set.mu.RLock()
	defer set.mu.RUnlock()
	buffer := bytes.NewBuffer(nil)
	l := len(set.data)
	i := 0
	for k, _ := range set.data {
		buffer.WriteString(`"` + gstr.QuoteMeta(k, `"\`) + `"`)
		if i != l-1 {
			buffer.WriteByte(',')
		}
		i++
	}
	return buffer.String()
}

// LockFunc locks writing with callback function <f>.
func (set *StrSet) LockFunc(f func(m map[string]struct{})) {
	set.mu.Lock()
	defer set.mu.Unlock()
	f(set.data)
}

// RLockFunc locks reading with callback function <f>.
func (set *StrSet) RLockFunc(f func(m map[string]struct{})) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	f(set.data)
}

// Equal checks whether the two sets equal.
func (set *StrSet) Equal(other *StrSet) bool {
	if set == other {
		return true
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	if len(set.data) != len(other.data) {
		return false
	}
	for key := range set.data {
		if _, ok := other.data[key]; !ok {
			return false
		}
	}
	return true
}

// IsSubsetOf checks whether the current set is a sub-set of <other>.
func (set *StrSet) IsSubsetOf(other *StrSet) bool {
	if set == other {
		return true
	}
	set.mu.RLock()
	defer set.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	for key := range set.data {
		if _, ok := other.data[key]; !ok {
			return false
		}
	}
	return true
}

// Union returns a new set which is the union of <set> and <other>.
// Which means, all the items in <newSet> are in <set> or in <other>.
func (set *StrSet) Union(others ...*StrSet) (newSet *StrSet) {
	newSet = NewStrSet(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range set.data {
			newSet.data[k] = v
		}
		if set != other {
			for k, v := range other.data {
				newSet.data[k] = v
			}
		}
		if set != other {
			other.mu.RUnlock()
		}
	}

	return
}

// Diff returns a new set which is the difference set from <set> to <other>.
// Which means, all the items in <newSet> are in <set> but not in <other>.
func (set *StrSet) Diff(others ...*StrSet) (newSet *StrSet) {
	newSet = NewStrSet(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set == other {
			continue
		}
		other.mu.RLock()
		for k, v := range set.data {
			if _, ok := other.data[k]; !ok {
				newSet.data[k] = v
			}
		}
		other.mu.RUnlock()
	}
	return
}

// Intersect returns a new set which is the intersection from <set> to <other>.
// Which means, all the items in <newSet> are in <set> and also in <other>.
func (set *StrSet) Intersect(others ...*StrSet) (newSet *StrSet) {
	newSet = NewStrSet(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	for _, other := range others {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range set.data {
			if _, ok := other.data[k]; ok {
				newSet.data[k] = v
			}
		}
		if set != other {
			other.mu.RUnlock()
		}
	}
	return
}

// Complement returns a new set which is the complement from <set> to <full>.
// Which means, all the items in <newSet> are in <full> and not in <set>.
//
// It returns the difference between <full> and <set>
// if the given set <full> is not the full set of <set>.
func (set *StrSet) Complement(full *StrSet) (newSet *StrSet) {
	newSet = NewStrSet(true)
	set.mu.RLock()
	defer set.mu.RUnlock()
	if set != full {
		full.mu.RLock()
		defer full.mu.RUnlock()
	}
	for k, v := range full.data {
		if _, ok := set.data[k]; !ok {
			newSet.data[k] = v
		}
	}
	return
}

// Merge adds items from <others> sets into <set>.
func (set *StrSet) Merge(others ...*StrSet) *StrSet {
	set.mu.Lock()
	defer set.mu.Unlock()
	for _, other := range others {
		if set != other {
			other.mu.RLock()
		}
		for k, v := range other.data {
			set.data[k] = v
		}
		if set != other {
			other.mu.RUnlock()
		}
	}
	return set
}

// Sum sums items.
// Note: The items should be converted to int type,
// or you'd get a result that you unexpected.
func (set *StrSet) Sum() (sum int) {
	set.mu.RLock()
	defer set.mu.RUnlock()
	for k, _ := range set.data {
		sum += gconv.Int(k)
	}
	return
}

// Pops randomly pops an item from set.
func (set *StrSet) Pop() string {
	set.mu.Lock()
	defer set.mu.Unlock()
	for k, _ := range set.data {
		delete(set.data, k)
		return k
	}
	return ""
}

// Pops randomly pops <size> items from set.
// It returns all items if size == -1.
func (set *StrSet) Pops(size int) []string {
	set.mu.Lock()
	defer set.mu.Unlock()
	if size > len(set.data) || size == -1 {
		size = len(set.data)
	}
	if size <= 0 {
		return nil
	}
	index := 0
	array := make([]string, size)
	for k, _ := range set.data {
		delete(set.data, k)
		array[index] = k
		index++
		if index == size {
			break
		}
	}
	return array
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (set *StrSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.Slice())
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (set *StrSet) UnmarshalJSON(b []byte) error {
	if set.mu == nil {
		set.mu = rwmutex.New()
		set.data = make(map[string]struct{})
	}
	set.mu.Lock()
	defer set.mu.Unlock()
	var array []string
	if err := json.Unmarshal(b, &array); err != nil {
		return err
	}
	for _, v := range array {
		set.data[v] = struct{}{}
	}
	return nil
}
