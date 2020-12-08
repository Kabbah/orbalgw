package mqttsn

import (
	"encoding/binary"
	"fmt"
)

// PublishMessage represents the contents of a MQTT-SN PUBLISH message.
type PublishMessage struct {
	Flags     Flags
	TopicID   uint16
	MessageID uint16
	Data      []byte
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *PublishMessage) MarshalBinary() ([]byte, error) {
	body := make([]byte, 5, 5+len(m.Data))
	flags, err := m.Flags.Value()
	body[0] = flags
	binary.BigEndian.PutUint16(body[1:3], m.TopicID)
	binary.BigEndian.PutUint16(body[3:5], m.MessageID)
	return append(body, m.Data...), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *PublishMessage) UnmarshalBinary(body []byte) error {
	if len(body) < 5 {
		return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}
	err := m.Flags.Parse(body[0])
	m.TopicID = binary.BigEndian.Uint16(body[1:3])
	m.MessageID = binary.BigEndian.Uint16(body[3:5])
	m.Data = make([]byte, len(body)-5)
	copy(m.Data, body[5:])
	return err
}
