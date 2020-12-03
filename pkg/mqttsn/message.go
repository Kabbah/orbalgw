package mqttsn

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// The MessageType field in MQTT-SN is 1-byte long.
type MessageType uint8

// Values of the MessageType field.
const (
	Advertise MessageType = iota
	SearchGw
	GwInfo
	_
	Connect
	ConnAck
	WillTopicReq
	WillTopic
	WillMsgReq
	WillMsg
	Register
	RegAck
	Publish
	PubAck
	PubComp
	PubRec
	PubRel
	_
	Subscribe
	SubAck
	Unsubscribe
	UnsubAck
	PingReq
	PingResp
	Disconnect
	_
	WillTopicUpd
	WillTopicResp
	WillMsgUpd
	WillMsgResp
	Encapsulated MessageType = 254
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
	length := len(m.Body)
	if length > maxBodyLength3Bytes {
		return nil, fmt.Errorf("message: length (%v) exceeds MQTT-SN limit", length)
	}

	var header []byte
	if length > maxBodyLength1Byte {
		length += 4
		header = make([]byte, 4, length)
		header[0] = 1
		binary.BigEndian.PutUint16(header[1:3], uint16(length))
		header[3] = uint8(m.Type)
	} else {
		length += 2
		header = make([]byte, 2, length)
		header[0] = uint8(length)
		header[1] = uint8(m.Type)
	}

	return append(header, m.Body...), nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.UnmarshalBinary by parsing the contents of a MQTT-SN binary
// message. The body (variable part) is just sliced from the original buffer as is, and should be parsed afterwards.
func (m *Message) UnmarshalBinary(data []byte) error {
	if len(data) < 2 {
		return errors.New("message: invalid MQTT-SN buffer")
	}

	var length int
	if data[0] == 1 {
		if len(data) < 4 {
			return errors.New("message: invalid MQTT-SN buffer")
		}
		length = int(binary.BigEndian.Uint16(data[1:3]))
		m.Type = MessageType(data[3])
		m.Body = data[4:]
	} else {
		length = int(data[0])
		m.Type = MessageType(data[1])
		m.Body = data[2:]
	}

	if len(data) != length {
		return errors.New("message: MQTT-SN buffer has incorrect length")
	}
	return nil
}
