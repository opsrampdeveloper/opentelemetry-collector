// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pdiscovery // import "go.opentelemetry.io/collector/pdata/plog"

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlpdisc "go.opentelemetry.io/collector/pdata/internal/data/protogen/discovery/v1"
)

//var _ MarshalSizer = (*ProtoMarshaler)(nil)

type ProtoMarshaler struct{}

func (e *ProtoMarshaler) MarshalDiscovery(ld Discovery) ([]byte, error) {
	pb := internal.DiscoveryToProto(internal.Discovery(ld))
	return pb.Marshal()
}

func (e *ProtoMarshaler) DiscoverySize(ld Discovery) int {
	pb := internal.DiscoveryToProto(internal.Discovery(ld))
	return pb.Size()
}

//var _ Unmarshaler = (*ProtoUnmarshaler)(nil)

type ProtoUnmarshaler struct{}

func (d *ProtoUnmarshaler) UnmarshalDiscovery(buf []byte) (Discovery, error) {
	pb := otlpdisc.DiscoveryData{}
	err := pb.Unmarshal(buf)
	return Discovery(internal.DiscoveryFromProto(pb)), err
}
