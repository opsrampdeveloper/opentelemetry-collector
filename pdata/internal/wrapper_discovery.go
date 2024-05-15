// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "go.opentelemetry.io/collector/pdata/internal"

import (
	otlpcollectordisc "go.opentelemetry.io/collector/pdata/internal/data/protogen/collector/discovery/v1"
	otlpdisc "go.opentelemetry.io/collector/pdata/internal/data/protogen/discovery/v1"
)

type Discovery struct {
	orig  *otlpcollectordisc.ExportDiscoveryServiceRequest
	state *State
}

func GetOrigDiscovery(ms Discovery) *otlpcollectordisc.ExportDiscoveryServiceRequest {
	return ms.orig
}

func GetDiscoveryState(ms Discovery) *State {
	return ms.state
}

func SetDiscoveryState(ms Discovery, state State) {
	*ms.state = state
}

func NewDiscovery(orig *otlpcollectordisc.ExportDiscoveryServiceRequest, state *State) Discovery {
	return Discovery{orig: orig, state: state}
}

// DiscoveryToProto internal helper to convert Discovery to protobuf representation.
func DiscoveryToProto(l Discovery) otlpdisc.DiscoveryData {
	return otlpdisc.DiscoveryData{
		ResourceDiscovery: l.orig.ResourceDiscovery,
	}
}

// DiscoveryFromProto internal helper to convert protobuf representation to Discovery.
// This function set exclusive state assuming that it's called only once per Discovery.
func DiscoveryFromProto(orig otlpdisc.DiscoveryData) Discovery {
	state := StateMutable
	return NewDiscovery(&otlpcollectordisc.ExportDiscoveryServiceRequest{
		ResourceDiscovery: orig.ResourceDiscovery,
	}, &state)
}
