package mqttsn

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestMessageMarshal(t *testing.T) {
	body := make([]byte, 65532)
	tests := []Message{
		Message{Advertise, body[:3]},
		Message{PingReq, nil},
		Message{Publish, body[:maxBodyLength1Byte]},
		Message{Publish, body[:maxBodyLength1Byte+1]},
		Message{Publish, body[:maxBodyLength3Bytes]},
		Message{Publish, body[:maxBodyLength3Bytes+1]},
		Message{Publish, body},
	}

	for _, msg := range tests {
		buf, err := msg.Marshal()
		if err == nil {
			headerLen := len(buf) - len(msg.Body)
			if len(msg.Body) > maxBodyLength3Bytes {
				t.Errorf("Expected error, but got nil, len: %v", len(msg.Body))
			} else if len(msg.Body) > maxBodyLength1Byte {
				length := int(binary.BigEndian.Uint16(buf[1:3]))
				if headerLen != 4 {
					t.Errorf("Expected 4-byte header, got %v", headerLen)
				}
				if buf[0] != 1 {
					t.Errorf("Expected buf to start with 0x01, got %v", buf[0])
				}
				if length != len(buf) {
					t.Errorf("Expected length field to be %v, got %v", len(buf), length)
				}
				if buf[3] != uint8(msg.Type) {
					t.Errorf("Expected message type field to be %v, got %v", uint8(msg.Type), buf[3])
				}
				if bytes.Compare(buf[4:], msg.Body) != 0 {
					t.Error("Message body is wrong")
				}
			} else {
				if headerLen != 2 {
					t.Errorf("Expected 2-byte header, got %v", headerLen)
				}
				if int(buf[0]) != len(buf) {
					t.Errorf("Expected length field to be %v, got %v", len(buf), buf[0])
				}
				if buf[1] != uint8(msg.Type) {
					t.Errorf("Expected message type field to be %v, got %v", uint8(msg.Type), buf[1])
				}
				if bytes.Compare(buf[2:], msg.Body) != 0 {
					t.Error("Message body is wrong")
				}
			}
		} else if len(msg.Body) <= maxBodyLength3Bytes {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
