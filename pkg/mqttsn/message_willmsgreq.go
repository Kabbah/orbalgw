package mqttsn

import "fmt"

// WillMsgReqMessage represents the contents of a MQTT-SN WILLMSGREQ message.
type WillMsgReqMessage struct{}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *WillMsgReqMessage) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *WillMsgReqMessage) UnmarshalBinary(body []byte) error {
	if len(body) != 0 {
		return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}
	return nil
}
