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
