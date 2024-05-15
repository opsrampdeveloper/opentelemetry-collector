// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "go.opentelemetry.io/collector/pdata/internal/cmd/pdatagen/internal"
import (
	"path/filepath"
)

var pdiscoveryotlp = &Package{
	name: "pdiscoveryotlp",
	path: filepath.Join("pdiscovery", "pdiscoveryotlp"),
	imports: []string{
		`otlpcollectordiscovery "go.opentelemetry.io/collector/pdata/internal/data/protogen/collector/discovery/v1"`,
	},
	testImports: []string{
		`"testing"`,
		``,
		`"github.com/stretchr/testify/assert"`,
	},
	structs: []baseStruct{
		exportDiscoveryPartialSuccess,
	},
}

var exportDiscoveryPartialSuccess = &messageValueStruct{
	structName:     "ExportPartialSuccess",
	description:    "// ExportPartialSuccess represents the details of a partially successful export request.",
	originFullName: "otlpcollectordiscovery.ExportDiscoveryPartialSuccess",
	fields: []baseField{
		&primitiveField{
			fieldName:  "RejectedDiscoveryRecords",
			returnType: "int64",
			defaultVal: `int64(0)`,
			testVal:    `int64(13)`,
		},
		&primitiveField{
			fieldName:  "ErrorMessage",
			returnType: "string",
			defaultVal: `""`,
			testVal:    `"error message"`,
		},
	},
}
