package mqttsn

import (
	"bytes"
	"testing"
)

func TestPingreqMarshal(t *testing.T) {
	tests := []struct {
		msg PingReqMessage
		buf []byte
	}{
		{PingReqMessage{}, []byte{}},
		{PingReqMessage{"0123"}, []byte{0x30, 0x31, 0x32, 0x33}},
	}

	for _, tt := range tests {
		if buf, err := tt.msg.MarshalBinary(); err == nil {
			if !bytes.Equal(buf, tt.buf) {
				t.Error("Message body is wrong")
			}
		} else {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestPingreqUnmarshal(t *testing.T) {
	tests := []struct {
		buf []byte
		msg PingReqMessage
	}{
		{[]byte{}, PingReqMessage{}},
		{[]byte{0x30, 0x31, 0x32, 0x33}, PingReqMessage{"0123"}},
	}

	for _, tt := range tests {
		var msg PingReqMessage
		if err := msg.UnmarshalBinary(tt.buf); err == nil {
			if msg.ClientID != tt.msg.ClientID {
				t.Errorf("Expected ClientID to be %v, got %v", tt.msg.ClientID, msg.ClientID)
			}
		} else {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
