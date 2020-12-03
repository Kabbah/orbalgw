package mqttsn

import (
	"encoding/binary"
	"fmt"
)

// ProtocolID of MQTT-SN is always 0x01.
const protocolID uint8 = 1

// ConnectMessage represents the contents of a MQTT-SN CONNECT message.
type ConnectMessage struct {
	Flags    Flags
	Duration uint16
	ClientID string
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *ConnectMessage) MarshalBinary() ([]byte, error) {
	body := make([]byte, 4, 4+len(m.ClientID))

	flags, err := m.Flags.Value()
	if err != nil {
		return nil, err
	}
	body[0] = flags
	body[1] = protocolID
	binary.BigEndian.PutUint16(body[2:4], m.Duration)

	return append(body, m.ClientID...), nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *ConnectMessage) UnmarshalBinary(body []byte) error {
	if len(body) < 4 {
		return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}

	err := m.Flags.Parse(body[0])
	if err != nil {
		return err
	}
	if body[1] != protocolID {
		return fmt.Errorf("mqttsn: invalid protocol ID (%v)", body[1])
	}
	m.Duration = binary.BigEndian.Uint16(body[2:4])
	m.ClientID = string(body[4:])
	return nil
}
