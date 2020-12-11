package mqttsn

// WillMsgMessage represents the contents of a MQTT-SN WILLMSG message.
type WillMsgMessage struct {
	Data []byte
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *WillMsgMessage) MarshalBinary() ([]byte, error) {
	body := make([]byte, len(m.Data))
	copy(body, m.Data)
	return body, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *WillMsgMessage) UnmarshalBinary(body []byte) error {
	m.Data = make([]byte, len(body))
	copy(m.Data, body)
	return nil
}
