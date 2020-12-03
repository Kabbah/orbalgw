package mqttsn

import "errors"

// SearchGwMessage represents the contents of a MQTT-SN SEARCHGW message.
type SearchGwMessage struct {
	Radius uint8
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *SearchGwMessage) MarshalBinary() ([]byte, error) {
	body := make([]byte, 1)
	body[0] = m.Radius
	return body, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *SearchGwMessage) UnmarshalBinary(body []byte) error {
	if len(body) != 1 {
		return errors.New("message: body has invalid size")
	}

	m.Radius = body[0]
	return nil
}
