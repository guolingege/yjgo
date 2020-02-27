// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package grand provides high performance random string generation functionality.
package grand

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	digits  = []rune("0123456789")
)

// Meet randomly calculate whether the given probability <num>/<total> is met.
func Meet(num, total int) bool {
	return Intn(total) < num
}

// MeetProb randomly calculate whether the given probability is met.
func MeetProb(prob float32) bool {
	return Intn(1e7) < int(prob*1e7)
}

// N returns a random int between min and max - [min, max].
// The <min> and <max> also support negative numbers.
func N(min, max int) int {
	if min >= max {
		return min
	}
	if min >= 0 {
		// Because Intn dose not support negative number,
		// so we should first shift the value to left,
		// then call Intn to produce the random number,
		// and finally shift the result back to right.
		return Intn(max-(min-0)+1) + (min - 0)
	}
	if min < 0 {
		// Because Intn dose not support negative number,
		// so we should first shift the value to right,
		// then call Intn to produce the random number,
		// and finally shift the result back to left.
		return Intn(max+(0-min)+1) - (0 - min)
	}
	return 0
}

// Str returns a random string which contains digits and letters, and its length is <n>.
func Str(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[Intn(62)]
	}
	return string(b)
}

// Digits returns a random string which contains only digits, and its length is <n>.
func Digits(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = digits[Intn(10)]
	}
	return string(b)
}

// Letters returns a random string which contains only letters, and its length is <n>.
func Letters(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[Intn(52)]
	}
	return string(b)
}

// Perm returns, as a slice of n int numbers, a pseudo-random permutation of the integers [0,n).
func Perm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		j := Intn(i + 1)
		m[i] = m[j]
		m[j] = i
	}
	return m
}
