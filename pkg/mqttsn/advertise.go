package mqttsn

import (
	"encoding/binary"
	"errors"
)

// AdvertiseMessage represents the contents of a MQTT-SN ADVERTISE message.
type AdvertiseMessage struct {
	GatewayID uint8
	Duration  uint16
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *AdvertiseMessage) MarshalBinary() ([]byte, error) {
	body := make([]byte, 3)
	body[0] = m.GatewayID
	binary.BigEndian.PutUint16(body[1:3], m.Duration)
	return body, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *AdvertiseMessage) UnmarshalBinary(body []byte) error {
	if len(body) != 3 {
		return errors.New("message: body has invalid size")
	}

	m.GatewayID = body[0]
	m.Duration = binary.BigEndian.Uint16(body[1:])
	return nil
}
