package mqttsn

import (
	"fmt"
)

// WillTopicReqMessage represents the contents of a MQTT-SN WILLTOPICREQ message.
type WillTopicReqMessage struct{}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *WillTopicReqMessage) MarshalBinary() ([]byte, error) {
	return []byte{}, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *WillTopicReqMessage) UnmarshalBinary(body []byte) error {
	if len(body) != 0 {
		return fmt.Errorf("mqttsn: invalid body length (%v)", len(body))
	}
	return nil
}
