// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

// MapPossibleItemByKey tries to find the possible key-value pair for given key ignoring cases and symbols.
//
// Note that this function might be of low performance.
func MapPossibleItemByKey(data map[string]interface{}, key string) (foundKey string, foundValue interface{}) {
	if len(data) == 0 {
		return
	}
	if v, ok := data[key]; ok {
		return key, v
	}
	// Loop checking.
	for k, v := range data {
		if EqualFoldWithoutChars(k, key) {
			return k, v
		}
	}
	return "", nil
}

// MapContainsPossibleKey checks if the given `key` is contained in given map `data`.
// It checks the key ignoring cases and symbols.
//
// Note that this function might be of low performance.
func MapContainsPossibleKey(data map[string]interface{}, key string) bool {
	if k, _ := MapPossibleItemByKey(data, key); k != "" {
		return true
	}
	return false
}
