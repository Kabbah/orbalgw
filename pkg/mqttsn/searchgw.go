package mqttsn

import "errors"

// SearchGwMessage represents the contents of a MQTT-SN SEARCHGW message.
type SearchGwMessage struct {
	Radius uint8
}

// Marshal converts a message to its binary form.
func (m *SearchGwMessage) Marshal() []byte {
	body := make([]byte, 1)
	body[0] = m.Radius
	return body
}

// Unmarshal parses a binary buffer into a message.
func (m *SearchGwMessage) Unmarshal(body []byte) error {
	if len(body) != 1 {
		return errors.New("message: body has invalid size")
	}

	m.Radius = body[0]
	return nil
}
