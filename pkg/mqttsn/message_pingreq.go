package mqttsn

// PingReqMessage represents the contents of a MQTT-SN PINGREQ message.
type PingReqMessage struct {
	ClientID string
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *PingReqMessage) MarshalBinary() ([]byte, error) {
	return []byte(m.ClientID), nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *PingReqMessage) UnmarshalBinary(body []byte) error {
	m.ClientID = string(body)
	return nil
}
