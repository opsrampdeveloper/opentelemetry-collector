// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pdiscovery // import "go.opentelemetry.io/collector/pdata/pdiscovery"

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"

	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/internal/json"
	"go.opentelemetry.io/collector/pdata/internal/otlp"
)

// JSONMarshaler marshals pdata.Discovery to JSON bytes using the OTLP/JSON format.
type JSONMarshaler struct{}

// MarshalDiscovery to the OTLP/JSON format.
func (*JSONMarshaler) MarshalDiscovery(ld Discovery) ([]byte, error) {
	buf := bytes.Buffer{}
	pb := internal.DiscoveryToProto(internal.Discovery(ld))
	err := json.Marshal(&buf, &pb)
	return buf.Bytes(), err
}

//var _ Unmarshaler = (*JSONUnmarshaler)(nil)

// JSONUnmarshaler unmarshals OTLP/JSON formatted-bytes to pdata.Discovery.
type JSONUnmarshaler struct{}

// UnmarshalDiscovery from OTLP/JSON format into pdata.Discovery.
func (*JSONUnmarshaler) UnmarshalDiscovery(buf []byte) (Discovery, error) {
	iter := jsoniter.ConfigFastest.BorrowIterator(buf)
	defer jsoniter.ConfigFastest.ReturnIterator(iter)
	ld := NewDiscovery()
	ld.unmarshalJsoniter(iter)
	if iter.Error != nil {
		return Discovery{}, iter.Error
	}
	otlp.MigrateDiscovery(ld.getOrig().ResourceDiscovery)
	return ld, nil
}

func (ms Discovery) unmarshalJsoniter(iter *jsoniter.Iterator) {
	iter.ReadObjectCB(func(iter *jsoniter.Iterator, f string) bool {
		switch f {
		case "resource_Discovery", "resourceDiscovery":
			iter.ReadArrayCB(func(*jsoniter.Iterator) bool {
				ms.ResourceDiscovery().AppendEmpty().unmarshalJsoniter(iter)
				return true
			})
		default:
			iter.Skip()
		}
		return true
	})
}

func (ms ResourceDiscovery) unmarshalJsoniter(iter *jsoniter.Iterator) {
	iter.ReadObjectCB(func(iter *jsoniter.Iterator, f string) bool {
		switch f {
		case "resource":
			json.ReadResource(iter, ms.orig.Resource)
		case "scope_discovery", "scopeDiscovery":
			iter.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
				ms.ScopeDiscovery().AppendEmpty().unmarshalJsoniter(iter)
				return true
			})
		case "schemaUrl", "schema_url":
			ms.orig.SchemaUrl = iter.ReadString()
		default:
			iter.Skip()
		}
		return true
	})
}

func (ms ScopeDiscovery) unmarshalJsoniter(iter *jsoniter.Iterator) {
	iter.ReadObjectCB(func(iter *jsoniter.Iterator, f string) bool {
		switch f {
		case "scope":
			json.ReadScope(iter, ms.orig.Scope)
		case "discovery_records", "discoveryRecords":
			iter.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
				ms.DiscoveryRecords().AppendEmpty().unmarshalJsoniter(iter)
				return true
			})
		case "schemaUrl", "schema_url":
			ms.orig.SchemaUrl = iter.ReadString()
		default:
			iter.Skip()
		}
		return true
	})
}

func (ms DiscoveryRecord) unmarshalJsoniter(iter *jsoniter.Iterator) {
	json.ReadResource(iter, ms.orig)
}
