// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pprofile // import "go.opentelemetry.io/collector/pdata/pprofile"

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlpprofile "go.opentelemetry.io/collector/pdata/internal/data/protogen/profiles/v1development"
)

var _ MarshalSizer = (*ProtoMarshaler)(nil)

type ProtoMarshaler struct{}

func (*ProtoMarshaler) MarshalProfiles(pd Profiles) ([]byte, error) {
	pb := internal.ProfilesToProto(internal.Profiles(pd))
	return pb.Marshal()
}

func (*ProtoMarshaler) ProfilesSize(pd Profiles) int {
	pb := internal.ProfilesToProto(internal.Profiles(pd))
	return pb.Size()
}

func (*ProtoMarshaler) ResourceProfilesSize(pd ResourceProfiles) int {
	return pd.orig.Size()
}

func (*ProtoMarshaler) ScopeProfilesSize(pd ScopeProfiles) int {
	return pd.orig.Size()
}

func (*ProtoMarshaler) ProfileSize(pd Profile) int {
	return pd.orig.Size()
}

type ProtoUnmarshaler struct{}

func (*ProtoUnmarshaler) UnmarshalProfiles(buf []byte) (Profiles, error) {
	pb := otlpprofile.ProfilesData{}
	err := pb.Unmarshal(buf)
	return Profiles(internal.ProfilesFromProto(pb)), err
}
