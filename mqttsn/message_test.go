package mqttsn

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestMessageMarshal(t *testing.T) {
	body := make([]byte, 65532)
	tests := []Message{
		{Advertise, body[:3]},
		{PingReq, nil},
		{Publish, body[:maxBodyLength1Byte]},
		{Publish, body[:maxBodyLength1Byte+1]},
		{Publish, body[:maxBodyLength3Bytes]},
		{Publish, body[:maxBodyLength3Bytes+1]},
		{Publish, body},
	}

	for _, msg := range tests {
		buf, err := msg.MarshalBinary()
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
				if !bytes.Equal(buf[4:], msg.Body) {
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
				if !bytes.Equal(buf[2:], msg.Body) {
					t.Error("Message body is wrong")
				}
			}
		} else if len(msg.Body) <= maxBodyLength3Bytes {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestMessageUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        Message
		shouldFail bool
	}{
		{buf: nil, shouldFail: true},
		{buf: []byte{}, shouldFail: true},
		{buf: []byte{0x02, 0x16}, msg: Message{PingReq, nil}},
		{buf: []byte{0x03, 0x16}, shouldFail: true},
		{buf: []byte{0x01, 0x00, 0x03}, shouldFail: true},
		{buf: []byte{0x01, 0x00, 0x04, 0x16}, msg: Message{PingReq, nil}},
		{buf: []byte{0x01, 0x00, 0x05, 0x16}, shouldFail: true},
	}

	for _, tt := range tests {
		var msg Message
		err := msg.UnmarshalBinary(tt.buf)
		if err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else {
				if msg.Type != tt.msg.Type {
					t.Errorf("Expected message type to be %v, got %v", msg.Type, tt.msg.Type)
				}
				if !bytes.Equal(msg.Body, tt.msg.Body) {
					t.Error("Message body is wrong")
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
