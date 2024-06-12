// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package str

import (
	"bytes"

	"github.com/gocarp/helpers/utils"
)

// AddSlashes quotes with slashes `\` for chars: '"\.
func AddSlashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// StripSlashes un-quotes a quoted string by AddSlashes.
func StripSlashes(str string) string {
	return utils.StripSlashes(str)
}

// QuoteMeta returns a version of `str` with a backslash character (`\`).
// If custom chars `chars` not given, it uses default chars: .\+*?[^]($)
func QuoteMeta(str string, chars ...string) string {
	var buf bytes.Buffer
	for _, char := range str {
		if len(chars) > 0 {
			for _, c := range chars[0] {
				if c == char {
					buf.WriteRune('\\')
					break
				}
			}
		} else {
			switch char {
			case '.', '+', '\\', '(', '$', ')', '[', '^', ']', '*', '?':
				buf.WriteRune('\\')
			}
		}
		buf.WriteRune(char)
	}
	return buf.String()
}
