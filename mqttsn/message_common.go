package mqttsn

import (
	"encoding/binary"
	"fmt"
)

type emptyMessage struct{}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *emptyMessage) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *emptyMessage) UnmarshalBinary(body []byte) error {
	if len(body) != 0 {
		return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}
	return nil
}

type messageIDOnlyMessage struct {
	MessageID uint16
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *messageIDOnlyMessage) MarshalBinary() ([]byte, error) {
	body := make([]byte, 2)
	binary.BigEndian.PutUint16(body[0:2], m.MessageID)
	return body, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *messageIDOnlyMessage) UnmarshalBinary(body []byte) error {
	if len(body) != 2 {
		return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}
	m.MessageID = binary.BigEndian.Uint16(body[0:2])
	return nil
}

type returnCodeOnlyMessage struct {
	ReturnCode ReturnCode
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *returnCodeOnlyMessage) MarshalBinary() ([]byte, error) {
	body := make([]byte, 1)
	body[0] = uint8(m.ReturnCode)
	return body, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *returnCodeOnlyMessage) UnmarshalBinary(body []byte) error {
	if len(body) != 1 {
		return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}

	m.ReturnCode = ReturnCode(body[0])
	return nil
}

type topicAckMessage struct {
	TopicID    uint16
	MessageID  uint16
	ReturnCode ReturnCode
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *topicAckMessage) MarshalBinary() ([]byte, error) {
	body := make([]byte, 5)
	binary.BigEndian.PutUint16(body[0:2], m.TopicID)
	binary.BigEndian.PutUint16(body[2:4], m.MessageID)
	body[4] = uint8(m.ReturnCode)
	return body, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *topicAckMessage) UnmarshalBinary(body []byte) error {
	if len(body) != 5 {
		return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}
	m.TopicID = binary.BigEndian.Uint16(body[0:2])
	m.MessageID = binary.BigEndian.Uint16(body[2:4])
	m.ReturnCode = ReturnCode(body[4])
	return nil
}
