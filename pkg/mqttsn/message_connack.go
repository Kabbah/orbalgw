package mqttsn

import (
	"fmt"
)

// ConnAckMessage represents the contents of a MQTT-SN CONNACK message.
type ConnAckMessage struct {
	ReturnCode ReturnCode
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *ConnAckMessage) MarshalBinary() ([]byte, error) {
	body := make([]byte, 1)
	body[0] = uint8(m.ReturnCode)
	return body, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *ConnAckMessage) UnmarshalBinary(body []byte) error {
	if len(body) != 1 {
		return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}

	m.ReturnCode = ReturnCode(body[0])
	return nil
}
