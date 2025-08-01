// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statustest // import "go.opentelemetry.io/collector/service/internal/status/statustest"

import (
	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/service/internal/status"
)

func NewNopStatusReporter() status.Reporter {
	return &nopStatusReporter{}
}

type nopStatusReporter struct{}

func (*nopStatusReporter) Ready() {}

func (*nopStatusReporter) ReportStatus(*componentstatus.InstanceID, *componentstatus.Event) {}

func (*nopStatusReporter) ReportOKIfStarting(*componentstatus.InstanceID) {}
