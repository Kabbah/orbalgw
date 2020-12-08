package mqttsn

import (
	"encoding/binary"
	"fmt"
)

// DisconnectMessage represents the contents of a MQTT-SN DISCONNECT message.
type DisconnectMessage struct {
	Duration *uint16
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *DisconnectMessage) MarshalBinary() ([]byte, error) {
	var body []byte
	if m.Duration != nil {
		body = make([]byte, 2)
		binary.BigEndian.PutUint16(body, *m.Duration)
	} else {
		body = []byte{}
	}
	return body, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *DisconnectMessage) UnmarshalBinary(body []byte) error {
	var err error = nil
	if len(body) == 0 {
		m.Duration = nil
	} else if len(body) == 2 {
		m.Duration = new(uint16)
		*m.Duration = binary.BigEndian.Uint16(body)
	} else {
		err = fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}
	return err
}
