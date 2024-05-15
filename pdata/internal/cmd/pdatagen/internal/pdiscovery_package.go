// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "go.opentelemetry.io/collector/pdata/internal/cmd/pdatagen/internal"

var pdiscovery = &Package{
	name: "pdiscovery",
	path: "pdiscovery",
	imports: []string{
		`"sort"`,
		``,
		`"go.opentelemetry.io/collector/pdata/internal"`,
		`"go.opentelemetry.io/collector/pdata/internal/data"`,
		`otlpdiscovery "go.opentelemetry.io/collector/pdata/internal/data/protogen/discovery/v1"`,
		`"go.opentelemetry.io/collector/pdata/pcommon"`,
	},
	testImports: []string{
		`"testing"`,
		`"unsafe"`,
		``,
		`"github.com/stretchr/testify/assert"`,
		``,
		`"go.opentelemetry.io/collector/pdata/internal"`,
		`"go.opentelemetry.io/collector/pdata/internal/data"`,
		`otlpdiscovery "go.opentelemetry.io/collector/pdata/internal/data/protogen/discovery/v1"`,
		`"go.opentelemetry.io/collector/pdata/pcommon"`,
	},
	structs: []baseStruct{
		resourceDiscoverySlice,
		resourceDiscovery,
		scopeDiscoverySlice,
		scopeDiscovery,
		discoveryRecordSlice,
		discoveryRecord,
	},
}

var resourceDiscoverySlice = &sliceOfPtrs{
	structName: "ResourceDiscoverySlice",
	element:    resourceDiscovery,
}

var resourceDiscovery = &messageValueStruct{
	structName:     "ResourceDiscovery",
	description:    "// ResourceDiscovery is a collection of discovery from a Resource.",
	originFullName: "otlpdiscovery.ResourceDiscovery",
	fields: []baseField{
		resourceField,
		schemaURLField,
		&sliceField{
			fieldName:   "ScopeDiscovery",
			returnSlice: scopeDiscoverySlice,
		},
	},
}

var scopeDiscoverySlice = &sliceOfPtrs{
	structName: "ScopeDiscoverySlice",
	element:    scopeDiscovery,
}

var scopeDiscovery = &messageValueStruct{
	structName:     "ScopeDiscovery",
	description:    "// ScopeDiscovery is a collection of discovery from a LibraryInstrumentation.",
	originFullName: "otlpdiscovery.ScopeDiscovery",
	fields: []baseField{
		scopeField,
		schemaURLField,
		&sliceField{
			fieldName:   "DiscoveryRecords",
			returnSlice: discoveryRecordSlice,
		},
	},
}

var discoveryRecordSlice = &sliceOfPtrs{
	structName: "DiscoveryRecordSlice",
	element:    discoveryRecord,
}

var discoveryRecord = &messageValueStruct{
	structName:     "DiscoveryRecord",
	description:    "// DiscoveryRecord are experimental implementation of OpenTelemetry Discovery Data Model.\n",
	originFullName: "otlpdiscovery.DiscoveryRecord",
	fields: []baseField{
		resourceField,
	},
}
