// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package str

import (
	"bytes"
	"strings"
	"unicode"
)

// Count counts the number of `substr` appears in `s`.
// It returns 0 if no `substr` found in `s`.
func Count(s, substr string) int {
	return strings.Count(s, substr)
}

// CountI counts the number of `substr` appears in `s`, case-insensitively.
// It returns 0 if no `substr` found in `s`.
func CountI(s, substr string) int {
	return strings.Count(ToLower(s), ToLower(substr))
}

// CountWords returns information about words' count used in a string.
// It considers parameter `str` as unicode string.
func CountWords(str string) map[string]int {
	m := make(map[string]int)
	buffer := bytes.NewBuffer(nil)
	for _, r := range []rune(str) {
		if unicode.IsSpace(r) {
			if buffer.Len() > 0 {
				m[buffer.String()]++
				buffer.Reset()
			}
		} else {
			buffer.WriteRune(r)
		}
	}
	if buffer.Len() > 0 {
		m[buffer.String()]++
	}
	return m
}

// CountChars returns information about chars' count used in a string.
// It considers parameter `str` as unicode string.
func CountChars(str string, noSpace ...bool) map[string]int {
	m := make(map[string]int)
	countSpace := true
	if len(noSpace) > 0 && noSpace[0] {
		countSpace = false
	}
	for _, r := range []rune(str) {
		if !countSpace && unicode.IsSpace(r) {
			continue
		}
		m[string(r)]++
	}
	return m
}
