package congestion

import "github.com/mutdroco/mpquic_arm_raspi/internal/protocol"

type connectionStats struct {
	slowstartPacketsLost protocol.PacketNumber
	slowstartBytesLost   protocol.ByteCount
}
