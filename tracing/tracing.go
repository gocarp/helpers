// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tracing provides some utility functions for tracing functionality.
package tracing

import (
	"math"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/gocarp/encoding/binary"
	"github.com/gocarp/go/container/types"
	"github.com/gocarp/utils/rand"
)

var (
	randomInitSequence = int32(rand.Intn(math.MaxInt32))
	sequence           = types.NewInt32(randomInitSequence)
)

// NewIDs creates and returns a new trace and span ID.
func NewIDs() (traceID trace.TraceID, spanID trace.SpanID) {
	return NewTraceID(), NewSpanID()
}

// NewTraceID creates and returns a trace ID.
func NewTraceID() (traceID trace.TraceID) {
	var (
		timestampNanoBytes = binary.EncodeInt64(time.Now().UnixNano())
		sequenceBytes      = binary.EncodeInt32(sequence.Add(1))
		randomBytes        = rand.B(4)
	)
	copy(traceID[:], timestampNanoBytes)
	copy(traceID[8:], sequenceBytes)
	copy(traceID[12:], randomBytes)
	return
}

// NewSpanID creates and returns a span ID.
func NewSpanID() (spanID trace.SpanID) {
	copy(spanID[:], binary.EncodeInt64(time.Now().UnixNano()/1e3))
	copy(spanID[4:], rand.B(4))
	return
}
