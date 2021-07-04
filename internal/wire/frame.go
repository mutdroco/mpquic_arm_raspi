package wire

import (
	"bytes"

	"github.com/mutdroco/mpquic_arm_raspi/internal/protocol"
)

// A Frame in QUIC
type Frame interface {
	Write(b *bytes.Buffer, version protocol.VersionNumber) error
	MinLength(version protocol.VersionNumber) (protocol.ByteCount, error)
}
