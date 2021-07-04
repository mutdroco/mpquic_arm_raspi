package utils

import "github.com/mutdroco/mpquic_arm_raspi/internal/protocol"

// ByteInterval is an interval from one ByteCount to the other
// +gen linkedlist
type ByteInterval struct {
	Start protocol.ByteCount
	End   protocol.ByteCount
}
