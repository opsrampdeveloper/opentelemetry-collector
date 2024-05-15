// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlp // import "go.opentelemetry.io/collector/pdata/internal/otlp"

import (
	otlpdisc "go.opentelemetry.io/collector/pdata/internal/data/protogen/discovery/v1"
)

// MigrateLogs implements any translation needed due to deprecation in OTLP logs protocol.
// Any plog.Unmarshaler implementation from OTLP (proto/json) MUST call this, and the gRPC Server implementation.
func MigrateDiscovery(rls []*otlpdisc.ResourceDiscovery) {
	for _, rl := range rls {
		if len(rl.ScopeDiscovery) == 0 {
			rl.ScopeDiscovery = nil
		}
		//rl.DeprecatedScopeLogs = nil
	}
}
