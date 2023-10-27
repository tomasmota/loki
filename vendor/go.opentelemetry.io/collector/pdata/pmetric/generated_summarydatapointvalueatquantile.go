// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
)

// SummaryDataPointValueAtQuantile is a quantile value within a Summary data point.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSummaryDataPointValueAtQuantile function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SummaryDataPointValueAtQuantile struct {
	orig *otlpmetrics.SummaryDataPoint_ValueAtQuantile
}

func newSummaryDataPointValueAtQuantile(orig *otlpmetrics.SummaryDataPoint_ValueAtQuantile) SummaryDataPointValueAtQuantile {
	return SummaryDataPointValueAtQuantile{orig}
}

// NewSummaryDataPointValueAtQuantile creates a new empty SummaryDataPointValueAtQuantile.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewSummaryDataPointValueAtQuantile() SummaryDataPointValueAtQuantile {
	return newSummaryDataPointValueAtQuantile(&otlpmetrics.SummaryDataPoint_ValueAtQuantile{})
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms SummaryDataPointValueAtQuantile) MoveTo(dest SummaryDataPointValueAtQuantile) {
	*dest.orig = *ms.orig
	*ms.orig = otlpmetrics.SummaryDataPoint_ValueAtQuantile{}
}

// Quantile returns the quantile associated with this SummaryDataPointValueAtQuantile.
func (ms SummaryDataPointValueAtQuantile) Quantile() float64 {
	return ms.orig.Quantile
}

// SetQuantile replaces the quantile associated with this SummaryDataPointValueAtQuantile.
func (ms SummaryDataPointValueAtQuantile) SetQuantile(v float64) {
	ms.orig.Quantile = v
}

// Value returns the value associated with this SummaryDataPointValueAtQuantile.
func (ms SummaryDataPointValueAtQuantile) Value() float64 {
	return ms.orig.Value
}

// SetValue replaces the value associated with this SummaryDataPointValueAtQuantile.
func (ms SummaryDataPointValueAtQuantile) SetValue(v float64) {
	ms.orig.Value = v
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms SummaryDataPointValueAtQuantile) CopyTo(dest SummaryDataPointValueAtQuantile) {
	dest.SetQuantile(ms.Quantile())
	dest.SetValue(ms.Value())
}
