package helpers

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimeToProto converts a time.Time pointer into a timestamppb.Timestamp pointer
func TimeToProto(t *time.Time) *timestamppb.Timestamp {
	var converted *timestamppb.Timestamp
	if t == nil {
		return nil
	}

	converted, err := ptypes.TimestampProto(*t)
	if err != nil {
		converted = nil
	}
	return converted
}

// ProtoTimeToTime converts a timestamppb.Timestamp pointer into a time.Time pointer
func ProtoTimeToTime(t *timestamppb.Timestamp) *time.Time {
	c, err := ptypes.Timestamp(t)
	converted := &c
	if err != nil {
		converted = (*time.Time)(nil)
	}
	return converted
}

// IsStringInSlice checks whether a slice contains a string exactly
func IsStringInSlice(s string, list []string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}

// IsStringSliceEqual tells whether slices a and b contain the same string elements.
// A nil argument is equivalent to an empty slice.
func IsStringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// IsUint64SliceEqual tells whether slices a and b contain the same uint64 elements.
// A nil argument is equivalent to an empty slice.
func IsUint64SliceEqual(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// CheckNil checks whether two interfaces are nil
func CheckNil(m1, m2 interface{}) (isNil bool, resolve bool) {
	// if both nil, return true and resolve
	if m1 == nil && m2 == nil {
		return true, true
	}
	// if one is nil and the other not, return false and resolve
	if (m1 == nil) != (m2 == nil) {
		return false, true
	}
	// both are not nil, return false and don't resolve
	return false, false
}

// TimeDiff gives the time between two time.Time
func TimeDiff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
