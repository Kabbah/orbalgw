package mqttsn

import (
	"bytes"
	"testing"
)

func TestConnAckMarshal(t *testing.T) {
	tests := []struct {
		msg ConnAckMessage
		buf []byte
	}{
		{ConnAckMessage{Accepted}, []byte{0x00}},
	}

	for _, tt := range tests {
		buf, err := tt.msg.MarshalBinary()
		if err == nil {
			if !bytes.Equal(buf, tt.buf) {
				t.Error("Message body is wrong")
			}
		} else {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestConnAckUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        ConnAckMessage
		shouldFail bool
	}{
		{buf: nil, shouldFail: true},
		{buf: []byte{}, shouldFail: true},
		{buf: []byte{0x00}, msg: ConnAckMessage{Accepted}},
		{buf: []byte{0x00, 0x00}, shouldFail: true},
	}

	for _, tt := range tests {
		var msg ConnAckMessage
		if err := msg.UnmarshalBinary(tt.buf); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else {
				if msg.ReturnCode != tt.msg.ReturnCode {
					t.Errorf("Expected ReturnCode to be %v, got %v", tt.msg.ReturnCode, msg.ReturnCode)
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
