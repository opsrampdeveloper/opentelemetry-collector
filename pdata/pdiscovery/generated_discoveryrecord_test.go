// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pdiscovery

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/internal/data"
	otlpdiscovery "go.opentelemetry.io/collector/pdata/internal/data/protogen/discovery/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
)


func TestDiscoveryRecord_MoveTo(t *testing.T) {
	ms := generateTestDiscoveryRecord()
	dest := NewDiscoveryRecord()
	ms.MoveTo(dest)
	assert.Equal(t, NewDiscoveryRecord(), ms)
	assert.Equal(t, generateTestDiscoveryRecord(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newDiscoveryRecord(&otlpdiscovery.DiscoveryRecord{}, &sharedState)) })
	assert.Panics(t, func() { newDiscoveryRecord(&otlpdiscovery.DiscoveryRecord{}, &sharedState).MoveTo(dest) })
}

func TestDiscoveryRecord_CopyTo(t *testing.T) {
	ms := NewDiscoveryRecord()
	orig := NewDiscoveryRecord()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestDiscoveryRecord()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newDiscoveryRecord(&otlpdiscovery.DiscoveryRecord{}, &sharedState)) })
}


func TestDiscoveryRecord_Resource(t *testing.T) {
	ms := NewDiscoveryRecord()
	internal.FillTestResource(internal.Resource(ms.Resource()))
	assert.Equal(t, pcommon.Resource(internal.GenerateTestResource()), ms.Resource())
}


func generateTestDiscoveryRecord() DiscoveryRecord {
	tv := NewDiscoveryRecord()
	fillTestDiscoveryRecord(tv)
	return tv
}

func fillTestDiscoveryRecord(tv DiscoveryRecord) {
	internal.FillTestResource(internal.NewResource(&tv.orig.Resource, tv.state))
}