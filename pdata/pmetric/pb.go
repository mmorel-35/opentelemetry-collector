// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pmetric // import "go.opentelemetry.io/collector/pdata/pmetric"

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
)

var _ MarshalSizer = (*ProtoMarshaler)(nil)

type ProtoMarshaler struct{}

func (*ProtoMarshaler) MarshalMetrics(md Metrics) ([]byte, error) {
	pb := internal.MetricsToProto(internal.Metrics(md))
	return pb.Marshal()
}

func (*ProtoMarshaler) MetricsSize(md Metrics) int {
	pb := internal.MetricsToProto(internal.Metrics(md))
	return pb.Size()
}

func (*ProtoMarshaler) ResourceMetricsSize(rm ResourceMetrics) int {
	return rm.orig.Size()
}

func (*ProtoMarshaler) ScopeMetricsSize(sm ScopeMetrics) int {
	return sm.orig.Size()
}

func (*ProtoMarshaler) MetricSize(m Metric) int {
	return m.orig.Size()
}

func (*ProtoMarshaler) NumberDataPointSize(ndp NumberDataPoint) int {
	return ndp.orig.Size()
}

func (*ProtoMarshaler) SummaryDataPointSize(sdps SummaryDataPoint) int {
	return sdps.orig.Size()
}

func (*ProtoMarshaler) HistogramDataPointSize(hdp HistogramDataPoint) int {
	return hdp.orig.Size()
}

func (*ProtoMarshaler) ExponentialHistogramDataPointSize(ehdp ExponentialHistogramDataPoint) int {
	return ehdp.orig.Size()
}

type ProtoUnmarshaler struct{}

func (*ProtoUnmarshaler) UnmarshalMetrics(buf []byte) (Metrics, error) {
	pb := otlpmetrics.MetricsData{}
	err := pb.Unmarshal(buf)
	return Metrics(internal.MetricsFromProto(pb)), err
}
