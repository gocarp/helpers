// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package str

import "strings"

// Repeat returns a new string consisting of multiplier copies of the string input.
//
// Example:
// Repeat("a", 3) -> "aaa"
func Repeat(input string, multiplier int) string {
	return strings.Repeat(input, multiplier)
}
