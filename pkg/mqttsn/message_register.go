package mqttsn

import (
	"encoding/binary"
	"fmt"
)

// RegisterMessage represents the contents of a MQTT-SN REGISTER message.
type RegisterMessage struct {
	TopicID   uint16
	MessageID uint16
	TopicName string
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *RegisterMessage) MarshalBinary() ([]byte, error) {
	topic := []byte(m.TopicName)
	body := make([]byte, 4, 4+len(topic))
	binary.BigEndian.PutUint16(body[0:2], m.TopicID)
	binary.BigEndian.PutUint16(body[2:4], m.MessageID)
	return append(body, topic...), nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *RegisterMessage) UnmarshalBinary(body []byte) error {
	if len(body) < 4 {
		return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}
	m.TopicID = binary.BigEndian.Uint16(body[0:2])
	m.MessageID = binary.BigEndian.Uint16(body[2:4])
	m.TopicName = string(body[4:])
	return nil
}
