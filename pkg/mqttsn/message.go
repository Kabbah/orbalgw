package mqttsn

import (
	"encoding/binary"
	"fmt"
)

// The MessageType field in MQTT-SN is 1-byte long.
type MessageType uint8

// Values of the MessageType field.
const (
	Advertise     MessageType = 0x00
	SearchGw      MessageType = 0x01
	GwInfo        MessageType = 0x02
	Connect       MessageType = 0x04
	ConnAck       MessageType = 0x05
	WillTopicReq  MessageType = 0x06
	WillTopic     MessageType = 0x07
	WillMsgReq    MessageType = 0x08
	WillMsg       MessageType = 0x09
	Register      MessageType = 0x0a
	RegAck        MessageType = 0x0b
	Publish       MessageType = 0x0c
	PubAck        MessageType = 0x0d
	PubComp       MessageType = 0x0e
	PubRec        MessageType = 0x0f
	PubRel        MessageType = 0x10
	Subscribe     MessageType = 0x12
	SubAck        MessageType = 0x13
	Unsubscribe   MessageType = 0x14
	UnsubAck      MessageType = 0x15
	PingReq       MessageType = 0x16
	PingResp      MessageType = 0x17
	Disconnect    MessageType = 0x18
	WillTopicUpd  MessageType = 0x1a
	WillTopicResp MessageType = 0x1b
	WillMsgUpd    MessageType = 0x1c
	WillMsgResp   MessageType = 0x1d
	Encapsulated  MessageType = 0xfe
)

const maxBodyLength1Byte int = 255 - 2
const maxBodyLength3Bytes int = 65535 - 4

// Message represents a full MQTT-SN message, which has a type and a body (variable part).
type Message struct {
	Type MessageType
	Body []byte
}

// MarshalBinary implements encoding.BinaryMarshaler.MarshalBinary by outputting a MQTT-SN binary message.
func (m *Message) MarshalBinary() ([]byte, error) {
	if len(m.Body) > maxBodyLength3Bytes {
		return nil, fmt.Errorf("mqttsn: body length (%v) exceeds MQTT-SN limit", len(m.Body))
	}

	var header []byte
	if len(m.Body) > maxBodyLength1Byte {
		length := 4 + len(m.Body)
		header = make([]byte, 4, length)
		header[0] = 1
		binary.BigEndian.PutUint16(header[1:3], uint16(length))
		header[3] = uint8(m.Type)
	} else {
		length := 2 + len(m.Body)
		header = make([]byte, 2, length)
		header[0] = uint8(length)
		header[1] = uint8(m.Type)
	}

	return append(header, m.Body...), nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary by parsing the contents of a MQTT-SN binary
// message. The body (variable part) is copied in binary form and should be parsed separately afterwards.
func (m *Message) UnmarshalBinary(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("mqttsn: invalid header size (%v)", len(data))
	}

	var length int
	var body []byte
	if data[0] == 1 {
		if len(data) < 4 {
			return fmt.Errorf("mqttsn: invalid header size (%v)", len(data))
		}
		length = int(binary.BigEndian.Uint16(data[1:3]))
		m.Type = MessageType(data[3])
		body = data[4:]
	} else {
		length = int(data[0])
		m.Type = MessageType(data[1])
		body = data[2:]
	}
	m.Body = make([]byte, len(body))
	copy(m.Body, body)

	if len(data) != length {
		return fmt.Errorf("mqttsn: header indicates incorrect message length (%v)", length)
	}
	return nil
}
