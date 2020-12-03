package mqttsn

import (
	"encoding/binary"
	"errors"
)

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
	body[1] = ProtocolID
	binary.BigEndian.PutUint16(body[2:4], m.Duration)

	return append(body, m.ClientID...), nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *ConnectMessage) UnmarshalBinary(body []byte) error {
	if len(body) < 4 {
		return errors.New("message: body has invalid size")
	}

	err := m.Flags.Parse(body[0])
	if err != nil {
		return err
	}
	if body[1] != ProtocolID {
		return errors.New("message: invalid protocol ID")
	}
	m.Duration = binary.BigEndian.Uint16(body[2:4])
	m.ClientID = string(body[4:])
	return nil
}
