package mqttsn

import "fmt"

// GwInfoMessage represents the contents of a MQTT-SN GWINFO message.
type GwInfoMessage struct {
	GatewayID      uint8
	GatewayAddress []byte
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *GwInfoMessage) MarshalBinary() ([]byte, error) {
	body := make([]byte, 1, 1+len(m.GatewayAddress))
	body[0] = m.GatewayID
	return append(body, m.GatewayAddress...), nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *GwInfoMessage) UnmarshalBinary(body []byte) error {
	if len(body) < 1 {
		return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}

	m.GatewayID = body[0]
	gwAddr := body[1:]
	m.GatewayAddress = make([]byte, len(gwAddr))
	copy(m.GatewayAddress, gwAddr)
	return nil
}
