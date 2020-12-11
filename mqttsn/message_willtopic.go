package mqttsn

import "errors"

// WillTopicMessage represents the contents of a MQTT-SN WILLTOPIC message.
type WillTopicMessage struct {
	Empty     bool
	Flags     Flags
	TopicName string
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *WillTopicMessage) MarshalBinary() ([]byte, error) {
	if m.Empty {
		return []byte{}, nil
	}
	flags, err := m.Flags.Value()
	return append([]byte{flags}, []byte(m.TopicName)...), err
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *WillTopicMessage) UnmarshalBinary(body []byte) error {
	if len(body) == 1 {
		return errors.New("mqttsn: invalid body length (1)")
	}

	if len(body) == 0 {
		m.Empty = true
		return nil
	}
	m.Empty = false
	err := m.Flags.Parse(body[0])
	m.TopicName = string(body[1:])
	return err
}
