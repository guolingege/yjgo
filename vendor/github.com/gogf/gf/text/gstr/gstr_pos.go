// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gstr

import "strings"

// Pos returns the position of the first occurrence of <needle>
// in <haystack> from <startOffset>, case-sensitively.
// It returns -1, if not found.
func Pos(haystack, needle string, startOffset ...int) int {
	length := len(haystack)
	offset := 0
	if len(startOffset) > 0 {
		offset = startOffset[0]
	}
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack[offset:], needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// PosI returns the position of the first occurrence of <needle>
// in <haystack> from <startOffset>, case-insensitively.
// It returns -1, if not found.
func PosI(haystack, needle string, startOffset ...int) int {
	length := len(haystack)
	offset := 0
	if len(startOffset) > 0 {
		offset = startOffset[0]
	}
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(strings.ToLower(haystack[offset:]), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// PosR returns the position of the last occurrence of <needle>
// in <haystack> from <startOffset>, case-sensitively.
// It returns -1, if not found.
func PosR(haystack, needle string, startOffset ...int) int {
	offset := 0
	if len(startOffset) > 0 {
		offset = startOffset[0]
	}
	pos, length := 0, len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		haystack = haystack[:offset+length+1]
	} else {
		haystack = haystack[offset:]
	}
	pos = strings.LastIndex(haystack, needle)
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}

// PosRI returns the position of the last occurrence of <needle>
// in <haystack> from <startOffset>, case-insensitively.
// It returns -1, if not found.
func PosRI(haystack, needle string, startOffset ...int) int {
	offset := 0
	if len(startOffset) > 0 {
		offset = startOffset[0]
	}
	pos, length := 0, len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		haystack = haystack[:offset+length+1]
	} else {
		haystack = haystack[offset:]
	}
	pos = strings.LastIndex(strings.ToLower(haystack), strings.ToLower(needle))
	if offset > 0 && pos != -1 {
		pos += offset
	}
	return pos
}
