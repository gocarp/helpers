// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package str

import "strings"

// IsSubDomain checks whether `subDomain` is subdomain of mainDomain.
// It supports '*' in `mainDomain`.
func IsSubDomain(subDomain string, mainDomain string) bool {
	if p := strings.IndexByte(subDomain, ':'); p != -1 {
		subDomain = subDomain[0:p]
	}
	if p := strings.IndexByte(mainDomain, ':'); p != -1 {
		mainDomain = mainDomain[0:p]
	}
	var (
		subArray   = strings.Split(subDomain, ".")
		mainArray  = strings.Split(mainDomain, ".")
		subLength  = len(subArray)
		mainLength = len(mainArray)
	)
	// Eg:
	// "gocarp.org" is not sub-domain of "s.gocarp.org".
	if mainLength > subLength {
		for i := range mainArray[0 : mainLength-subLength] {
			if mainArray[i] != "*" {
				return false
			}
		}
	}

	// Eg:
	// "s.s.gocarp.org" is not subdomain of "*.gocarp.org"
	// but
	// "s.s.gocarp.org" is sub-domain of "gocarp.org"
	if mainLength > 2 && subLength > mainLength {
		return false
	}
	minLength := subLength
	if mainLength < minLength {
		minLength = mainLength
	}
	for i := minLength; i > 0; i-- {
		if mainArray[mainLength-i] == "*" {
			continue
		}
		if mainArray[mainLength-i] != subArray[subLength-i] {
			return false
		}
	}
	return true
}
