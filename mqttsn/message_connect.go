package mqttsn

import (
	"encoding/binary"
	"fmt"
	"time"
)

// ProtocolID of MQTT-SN is always 0x01.
const protocolID uint8 = 1

// ConnectMessage represents the contents of a MQTT-SN CONNECT message.
type ConnectMessage struct {
	Flags    Flags
	Duration time.Duration
	ClientID string
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *ConnectMessage) MarshalBinary() ([]byte, error) {
	duration := m.Duration.Milliseconds() / 1000
	if duration < 0 || duration > 65535 {
		return nil, fmt.Errorf("mqttsn: Duration out of range (%v)", duration)
	}

	body := make([]byte, 4, 4+len(m.ClientID))
	flags, err := m.Flags.Value()
	body[0] = flags
	body[1] = protocolID
	binary.BigEndian.PutUint16(body[2:4], uint16(duration))

	return append(body, m.ClientID...), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *ConnectMessage) UnmarshalBinary(body []byte) error {
	if len(body) < 4 {
		return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}
	if body[1] != protocolID {
		return fmt.Errorf("mqttsn: invalid protocol ID (%v)", body[1])
	}

	err := m.Flags.Parse(body[0])
	duration := binary.BigEndian.Uint16(body[2:4])
	m.Duration = time.Duration(duration) * time.Second
	m.ClientID = string(body[4:])
	return err
}
