package wire

import "github.com/mutdroco/mpquic_arm_raspi/internal/protocol"

// AckRange is an ACK range
type AckRange struct {
	First protocol.PacketNumber
	Last  protocol.PacketNumber
}
