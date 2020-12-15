package mqttsn

import (
	"encoding/binary"
	"fmt"
	"time"
)

// AdvertiseMessage represents the contents of a MQTT-SN ADVERTISE message.
type AdvertiseMessage struct {
	GatewayID uint8
	Duration  time.Duration
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *AdvertiseMessage) MarshalBinary() ([]byte, error) {
	duration := m.Duration.Milliseconds() / 1000
	if duration < 0 || duration > 65535 {
		return nil, fmt.Errorf("mqttsn: Duration out of range (%v)", duration)
	}

	body := make([]byte, 3)
	body[0] = m.GatewayID
	binary.BigEndian.PutUint16(body[1:3], uint16(duration))
	return body, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *AdvertiseMessage) UnmarshalBinary(body []byte) error {
	if len(body) != 3 {
		return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}

	m.GatewayID = body[0]
	duration := binary.BigEndian.Uint16(body[1:3])
	m.Duration = time.Duration(duration) * time.Second
	return nil
}
