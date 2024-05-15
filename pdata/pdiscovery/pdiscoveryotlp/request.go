// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pdiscoveryotlp // import "go.opentelemetry.io/collector/pdata/plog/plogotlp"

import (
	"bytes"
	"go.opentelemetry.io/collector/pdata/internal"
	otlpcollectordisc "go.opentelemetry.io/collector/pdata/internal/data/protogen/collector/discovery/v1"
	"go.opentelemetry.io/collector/pdata/internal/json"
	"go.opentelemetry.io/collector/pdata/internal/otlp"
	"go.opentelemetry.io/collector/pdata/pdiscovery"
)

var jsonUnmarshaler = &pdiscovery.JSONUnmarshaler{}

// ExportRequest represents the request for gRPC/HTTP client/server.
// It's a wrapper for plog.Logs data.
type ExportRequest struct {
	orig  *otlpcollectordisc.ExportDiscoveryServiceRequest
	state *internal.State
}

// NewExportRequest returns an empty ExportRequest.
func NewExportRequest() ExportRequest {
	state := internal.StateMutable
	return ExportRequest{
		orig:  &otlpcollectordisc.ExportDiscoveryServiceRequest{},
		state: &state,
	}
}

// NewExportRequestFromLogs returns a ExportRequest from plog.Logs.
// Because ExportRequest is a wrapper for plog.Logs,
// any changes to the provided Logs struct will be reflected in the ExportRequest and vice versa.
func NewExportRequestFromDiscovery(ld pdiscovery.Discovery) ExportRequest {
	return ExportRequest{
		orig:  internal.GetOrigDiscovery(internal.Discovery(ld)),
		state: internal.GetDiscoveryState(internal.Discovery(ld)),
	}
}

// MarshalProto marshals ExportRequest into proto bytes.
func (ms ExportRequest) MarshalProto() ([]byte, error) {
	return ms.orig.Marshal()
}

// UnmarshalProto unmarshalls ExportRequest from proto bytes.
func (ms ExportRequest) UnmarshalProto(data []byte) error {
	if err := ms.orig.Unmarshal(data); err != nil {
		return err
	}
	otlp.MigrateDiscovery(ms.orig.ResourceDiscovery)
	return nil
}

// MarshalJSON marshals ExportRequest into JSON bytes.
func (ms ExportRequest) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if err := json.Marshal(&buf, ms.orig); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalJSON unmarshalls ExportRequest from JSON bytes.
func (ms ExportRequest) UnmarshalJSON(data []byte) error {
	ld, err := jsonUnmarshaler.UnmarshalDiscovery(data)
	if err != nil {
		return err
	}
	*ms.orig = *internal.GetOrigDiscovery(internal.Discovery(ld))
	return nil
}

func (ms ExportRequest) Discovery() pdiscovery.Discovery {
	return pdiscovery.Discovery(internal.NewDiscovery(ms.orig, ms.state))
}
