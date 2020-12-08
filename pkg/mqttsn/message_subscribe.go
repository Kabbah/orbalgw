package mqttsn

import (
	"encoding/binary"
	"fmt"
)

// SubscribeMessage represents the contents of a MQTT-SN SUBSCRIBE message.
type SubscribeMessage struct {
	Flags     Flags
	MessageID uint16
	TopicID   uint16
	TopicName string
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *SubscribeMessage) MarshalBinary() ([]byte, error) {
	var body []byte
	if m.Flags.TopicType == PredefinedTopicID {
		body = make([]byte, 5)
		binary.BigEndian.PutUint16(body[3:5], m.TopicID)
	} else {
		topic := []byte(m.TopicName)
		if m.Flags.TopicType == ShortTopicName && len(topic) != 2 {
			return nil, fmt.Errorf("mqttsn: invalid short topic name length (%v)", len(topic))
		}
		body = make([]byte, 3, 3+len(topic))
		body = append(body, topic...)
	}

	flags, err := m.Flags.Value()
	body[0] = flags
	binary.BigEndian.PutUint16(body[1:3], m.MessageID)
	return body, err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *SubscribeMessage) UnmarshalBinary(body []byte) error {
	if len(body) < 3 {
		return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}

	err := m.Flags.Parse(body[0])
	m.MessageID = binary.BigEndian.Uint16(body[1:3])
	if err != nil {
		return err
	}

	if m.Flags.TopicType == NormalTopicName {
		m.TopicName = string(body[3:])
	} else {
		if len(body) != 5 {
			return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
		}
		if m.Flags.TopicType == PredefinedTopicID {
			m.TopicID = binary.BigEndian.Uint16(body[3:5])
		} else {
			m.TopicName = string(body[3:5])
		}
	}
	return nil
}
