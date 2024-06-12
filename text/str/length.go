// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package str

import "unicode/utf8"

// LenRune returns string length of unicode.
func LenRune(str string) int {
	return utf8.RuneCountInString(str)
}
