// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package str

// SearchArray searches string `s` in string slice `a` case-sensitively,
// returns its index in `a`.
// If `s` is not found in `a`, it returns -1.
func SearchArray(a []string, s string) int {
	for i, v := range a {
		if s == v {
			return i
		}
	}
	return NotFoundIndex
}

// InArray checks whether string `s` in slice `a`.
func InArray(a []string, s string) bool {
	return SearchArray(a, s) != NotFoundIndex
}

// PrefixArray adds `prefix` string for each item of `array`.
//
// Example:
// PrefixArray(["a","b"], "gc_") -> ["gc_a", "gc_b"]
func PrefixArray(array []string, prefix string) {
	for k, v := range array {
		array[k] = prefix + v
	}
}
