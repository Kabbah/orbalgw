package mqttsn

import "fmt"

// EncapsulatedMessage represents an encapsulated MQTT-SN message, according to the Forwarder Encapsulation scheme.
type EncapsulatedMessage struct {
	Radius         uint8
	WirelessNodeID []byte
	Message
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary.
func (m *EncapsulatedMessage) MarshalBinary() ([]byte, error) {
	if m.Radius > 3 {
		return nil, fmt.Errorf("mqttsn: Radius too big (%v)", m.Radius)
	}
	if len(m.WirelessNodeID) > 255-3 {
		return nil, fmt.Errorf("mqttsn: WirelessNodeID too big (%v)", len(m.WirelessNodeID))
	}

	msg, err := m.Message.MarshalBinary()
	if err != nil {
		return nil, err
	}

	headerLen := 3 + len(m.WirelessNodeID)
	frame := make([]byte, (headerLen + len(msg)))
	frame[0] = uint8(headerLen)
	frame[1] = uint8(Encapsulated)
	frame[2] = m.Radius
	copy(frame[3:], m.WirelessNodeID)
	copy(frame[headerLen:], msg)
	return frame, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary.
func (m *EncapsulatedMessage) UnmarshalBinary(data []byte) error {
	if len(data) < 3 {
		return fmt.Errorf("mqttsn: invalid header length (%v)", len(data))
	}
	if MessageType(data[1]) != Encapsulated {
		return fmt.Errorf("mqttsn: type is incorrect (%v)", data[1])
	}

	headerLen := int(data[0])
	if len(data) < headerLen {
		return fmt.Errorf("mqttsn: header indicates incorrect frame length (%v)", headerLen)
	}

	m.Radius = data[2] & 0b0000_0011
	m.WirelessNodeID = make([]byte, (headerLen - 3))
	copy(m.WirelessNodeID, data[3:headerLen])
	return m.Message.UnmarshalBinary(data[headerLen:])
}
