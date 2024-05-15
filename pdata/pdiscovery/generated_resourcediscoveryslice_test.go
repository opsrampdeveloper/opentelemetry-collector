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

func TestResourceDiscoverySlice(t *testing.T) {
	es := NewResourceDiscoverySlice()
	assert.Equal(t, 0, es.Len())
	state := internal.StateMutable
	es = newResourceDiscoverySlice(&[]*otlpdiscovery.ResourceDiscovery{}, &state)
	assert.Equal(t, 0, es.Len())

	emptyVal := NewResourceDiscovery()
	testVal := generateTestResourceDiscovery()
	for i := 0; i < 7; i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, es.At(i))
		fillTestResourceDiscovery(el)
		assert.Equal(t, testVal, es.At(i))
	}
	assert.Equal(t, 7, es.Len())
}

func TestResourceDiscoverySliceReadOnly(t *testing.T) {
	sharedState := internal.StateReadOnly
	es := newResourceDiscoverySlice(&[]*otlpdiscovery.ResourceDiscovery{}, &sharedState)
	assert.Equal(t, 0, es.Len())
	assert.Panics(t, func() { es.AppendEmpty() })
	assert.Panics(t, func() { es.EnsureCapacity(2) })
	es2 := NewResourceDiscoverySlice()
	es.CopyTo(es2)
	assert.Panics(t, func() { es2.CopyTo(es) })
	assert.Panics(t, func() { es.MoveAndAppendTo(es2) })
	assert.Panics(t, func() { es2.MoveAndAppendTo(es) })
}

func TestResourceDiscoverySlice_CopyTo(t *testing.T) {
	dest := NewResourceDiscoverySlice()
	// Test CopyTo to empty
	NewResourceDiscoverySlice().CopyTo(dest)
	assert.Equal(t, NewResourceDiscoverySlice(), dest)

	// Test CopyTo larger slice
	generateTestResourceDiscoverySlice().CopyTo(dest)
	assert.Equal(t, generateTestResourceDiscoverySlice(), dest)

	// Test CopyTo same size slice
	generateTestResourceDiscoverySlice().CopyTo(dest)
	assert.Equal(t, generateTestResourceDiscoverySlice(), dest)
}

func TestResourceDiscoverySlice_EnsureCapacity(t *testing.T) {
	es := generateTestResourceDiscoverySlice()

	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	assert.Equal(t, es.Len(), cap(*es.orig))
	assert.Equal(t, generateTestResourceDiscoverySlice(), es)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	es.EnsureCapacity(ensureLargeLen)
	assert.Less(t, generateTestResourceDiscoverySlice().Len(), ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	assert.Equal(t, generateTestResourceDiscoverySlice(), es)
}

func TestResourceDiscoverySlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestResourceDiscoverySlice()
	dest := NewResourceDiscoverySlice()
	src := generateTestResourceDiscoverySlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestResourceDiscoverySlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestResourceDiscoverySlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestResourceDiscoverySlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestResourceDiscoverySlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewResourceDiscoverySlice()
	emptySlice.RemoveIf(func(el ResourceDiscovery) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestResourceDiscoverySlice()
	pos := 0
	filtered.RemoveIf(func(el ResourceDiscovery) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestResourceDiscoverySlice_Sort(t *testing.T) {
	es := generateTestResourceDiscoverySlice()
	es.Sort(func(a, b ResourceDiscovery) bool {
		return uintptr(unsafe.Pointer(a.orig)) < uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.True(t, uintptr(unsafe.Pointer(es.At(i-1).orig)) < uintptr(unsafe.Pointer(es.At(i).orig)))
	}
	es.Sort(func(a, b ResourceDiscovery) bool {
		return uintptr(unsafe.Pointer(a.orig)) > uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.True(t, uintptr(unsafe.Pointer(es.At(i-1).orig)) > uintptr(unsafe.Pointer(es.At(i).orig)))
	}
}

func generateTestResourceDiscoverySlice() ResourceDiscoverySlice {
	es := NewResourceDiscoverySlice()
	fillTestResourceDiscoverySlice(es)
	return es
}

func fillTestResourceDiscoverySlice(es ResourceDiscoverySlice) {
	*es.orig = make([]*otlpdiscovery.ResourceDiscovery, 7)
	for i := 0; i < 7; i++ {
		(*es.orig)[i] = &otlpdiscovery.ResourceDiscovery{}
		fillTestResourceDiscovery(newResourceDiscovery((*es.orig)[i], es.state))
	}
}