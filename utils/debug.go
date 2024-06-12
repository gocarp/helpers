// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import "github.com/gocarp/helpers/cmd"

const (
	// Debug key for checking if in debug mode.
	cmdEnvKeyForDebugKey = "gc.debug"
)

var (
	// isDebugEnabled marks whether GoCarp debug mode is enabled.
	isDebugEnabled = false
)

func init() {
	// Debugging configured.
	value := cmd.GetOptWithEnv(cmdEnvKeyForDebugKey)
	if value == "" || value == "0" || value == "false" {
		isDebugEnabled = false
	} else {
		isDebugEnabled = true
	}
}

// IsDebugEnabled checks and returns whether debug mode is enabled.
// The debug mode is enabled when cmd argument "gc.debug" or environment "GC_DEBUG" is passed.
func IsDebugEnabled() bool {
	return isDebugEnabled
}

// SetDebugEnabled enables/disables the internal debug info.
func SetDebugEnabled(enabled bool) {
	isDebugEnabled = enabled
}
