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

func TestScopeDiscoverySlice(t *testing.T) {
	es := NewScopeDiscoverySlice()
	assert.Equal(t, 0, es.Len())
	state := internal.StateMutable
	es = newScopeDiscoverySlice(&[]*otlpdiscovery.ScopeDiscovery{}, &state)
	assert.Equal(t, 0, es.Len())

	emptyVal := NewScopeDiscovery()
	testVal := generateTestScopeDiscovery()
	for i := 0; i < 7; i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, es.At(i))
		fillTestScopeDiscovery(el)
		assert.Equal(t, testVal, es.At(i))
	}
	assert.Equal(t, 7, es.Len())
}

func TestScopeDiscoverySliceReadOnly(t *testing.T) {
	sharedState := internal.StateReadOnly
	es := newScopeDiscoverySlice(&[]*otlpdiscovery.ScopeDiscovery{}, &sharedState)
	assert.Equal(t, 0, es.Len())
	assert.Panics(t, func() { es.AppendEmpty() })
	assert.Panics(t, func() { es.EnsureCapacity(2) })
	es2 := NewScopeDiscoverySlice()
	es.CopyTo(es2)
	assert.Panics(t, func() { es2.CopyTo(es) })
	assert.Panics(t, func() { es.MoveAndAppendTo(es2) })
	assert.Panics(t, func() { es2.MoveAndAppendTo(es) })
}

func TestScopeDiscoverySlice_CopyTo(t *testing.T) {
	dest := NewScopeDiscoverySlice()
	// Test CopyTo to empty
	NewScopeDiscoverySlice().CopyTo(dest)
	assert.Equal(t, NewScopeDiscoverySlice(), dest)

	// Test CopyTo larger slice
	generateTestScopeDiscoverySlice().CopyTo(dest)
	assert.Equal(t, generateTestScopeDiscoverySlice(), dest)

	// Test CopyTo same size slice
	generateTestScopeDiscoverySlice().CopyTo(dest)
	assert.Equal(t, generateTestScopeDiscoverySlice(), dest)
}

func TestScopeDiscoverySlice_EnsureCapacity(t *testing.T) {
	es := generateTestScopeDiscoverySlice()

	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	assert.Equal(t, es.Len(), cap(*es.orig))
	assert.Equal(t, generateTestScopeDiscoverySlice(), es)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	es.EnsureCapacity(ensureLargeLen)
	assert.Less(t, generateTestScopeDiscoverySlice().Len(), ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	assert.Equal(t, generateTestScopeDiscoverySlice(), es)
}

func TestScopeDiscoverySlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestScopeDiscoverySlice()
	dest := NewScopeDiscoverySlice()
	src := generateTestScopeDiscoverySlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestScopeDiscoverySlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestScopeDiscoverySlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestScopeDiscoverySlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestScopeDiscoverySlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewScopeDiscoverySlice()
	emptySlice.RemoveIf(func(el ScopeDiscovery) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestScopeDiscoverySlice()
	pos := 0
	filtered.RemoveIf(func(el ScopeDiscovery) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestScopeDiscoverySlice_Sort(t *testing.T) {
	es := generateTestScopeDiscoverySlice()
	es.Sort(func(a, b ScopeDiscovery) bool {
		return uintptr(unsafe.Pointer(a.orig)) < uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.True(t, uintptr(unsafe.Pointer(es.At(i-1).orig)) < uintptr(unsafe.Pointer(es.At(i).orig)))
	}
	es.Sort(func(a, b ScopeDiscovery) bool {
		return uintptr(unsafe.Pointer(a.orig)) > uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.True(t, uintptr(unsafe.Pointer(es.At(i-1).orig)) > uintptr(unsafe.Pointer(es.At(i).orig)))
	}
}

func generateTestScopeDiscoverySlice() ScopeDiscoverySlice {
	es := NewScopeDiscoverySlice()
	fillTestScopeDiscoverySlice(es)
	return es
}

func fillTestScopeDiscoverySlice(es ScopeDiscoverySlice) {
	*es.orig = make([]*otlpdiscovery.ScopeDiscovery, 7)
	for i := 0; i < 7; i++ {
		(*es.orig)[i] = &otlpdiscovery.ScopeDiscovery{}
		fillTestScopeDiscovery(newScopeDiscovery((*es.orig)[i], es.state))
	}
}