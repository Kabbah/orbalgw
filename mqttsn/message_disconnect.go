package mqttsn

import (
	"encoding/binary"
	"fmt"
	"time"
)

// DisconnectMessage represents the contents of a MQTT-SN DISCONNECT message.
type DisconnectMessage struct {
	Duration *time.Duration
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *DisconnectMessage) MarshalBinary() ([]byte, error) {
	var body []byte
	if m.Duration != nil {
		duration := m.Duration.Milliseconds() / 1000
		if duration < 0 || duration > 65535 {
			return nil, fmt.Errorf("mqttsn: Duration out of range (%v)", duration)
		}

		body = make([]byte, 2)
		binary.BigEndian.PutUint16(body, uint16(duration))
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
		duration := binary.BigEndian.Uint16(body)
		m.Duration = new(time.Duration)
		*m.Duration = time.Duration(duration) * time.Second
	} else {
		err = fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}
	return err
}
