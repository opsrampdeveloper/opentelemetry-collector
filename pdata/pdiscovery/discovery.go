// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pdiscovery // import "go.opentelemetry.io/collector/pdata/plog"

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlpcollectordisc "go.opentelemetry.io/collector/pdata/internal/data/protogen/collector/discovery/v1"
)

// Discovery is the top-level struct that is propagated through the logs pipeline.
// Use NewDiscovery to create new instance, zero-initialized instance is not valid for use.
type Discovery internal.Discovery

func newDiscovery(orig *otlpcollectordisc.ExportDiscoveryServiceRequest) Discovery {
	state := internal.StateMutable
	return Discovery(internal.NewDiscovery(orig, &state))
}

func (ms Discovery) getOrig() *otlpcollectordisc.ExportDiscoveryServiceRequest {
	return internal.GetOrigDiscovery(internal.Discovery(ms))
}

func (ms Discovery) getState() *internal.State {
	return internal.GetDiscoveryState(internal.Discovery(ms))
}

// NewDiscovery creates a new Discovery struct.
func NewDiscovery() Discovery {
	return newDiscovery(&otlpcollectordisc.ExportDiscoveryServiceRequest{})
}

// IsReadOnly returns true if this Discovery instance is read-only.
func (ms Discovery) IsReadOnly() bool {
	return *ms.getState() == internal.StateReadOnly
}

// CopyTo copies the Discovery instance overriding the destination.
func (ms Discovery) CopyTo(dest Discovery) {
	ms.ResourceDiscovery().CopyTo(dest.ResourceDiscovery())
}

// DiscoveryRecordCount calculates the total number of discovery records.
func (ms Discovery) DiscoveryRecordCount() int {
	discoveryCount := 0
	rss := ms.ResourceDiscovery()
	for i := 0; i < rss.Len(); i++ {
		rs := rss.At(i)
		ill := rs.ScopeDiscovery()
		for i := 0; i < ill.Len(); i++ {
			discovery := ill.At(i)
			discoveryCount += discovery.DiscoveryRecords().Len()
		}
	}
	return discoveryCount
}

// ResourceDiscovery returns the ResourceDiscoverySlice associated with this Discovery.
func (ms Discovery) ResourceDiscovery() ResourceDiscoverySlice {
	return newResourceDiscoverySlice(&ms.getOrig().ResourceDiscovery, internal.GetDiscoveryState(internal.Discovery(ms)))
}

// MarkReadOnly marks the Discovery as shared so that no further modifications can be done on it.
func (ms Discovery) MarkReadOnly() {
	internal.SetDiscoveryState(internal.Discovery(ms), internal.StateReadOnly)
}
