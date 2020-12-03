package mqttsn

import "errors"

// GwInfoMessage represents the contents of a MQTT-SN GWINFO message.
type GwInfoMessage struct {
	GatewayID      uint8
	GatewayAddress []byte
}

// Marshal converts a message to its binary form.
func (m *GwInfoMessage) Marshal() ([]byte, error) {
	body := make([]byte, 1, 1+len(m.GatewayAddress))
	body[0] = m.GatewayID
	return append(body, m.GatewayAddress...), nil
}

// Unmarshal parses a binary buffer into a message.
func (m *GwInfoMessage) Unmarshal(body []byte) error {
	if len(body) < 1 {
		return errors.New("message: body has invalid size")
	}

	m.GatewayID = body[0]
	m.GatewayAddress = body[1:]
	return nil
}
