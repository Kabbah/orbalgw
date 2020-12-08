package mqttsn

import (
	"bytes"
	"testing"
)

func TestDisconnectMarshal(t *testing.T) {
	var duration uint16 = 10
	tests := []struct {
		msg DisconnectMessage
		buf []byte
	}{
		{DisconnectMessage{nil}, []byte{}},
		{DisconnectMessage{&duration}, []byte{0x00, 0x0a}},
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

func TestDisconnectUnmarshal(t *testing.T) {
	var duration uint16 = 10
	tests := []struct {
		buf        []byte
		msg        DisconnectMessage
		shouldFail bool
	}{
		{buf: []byte{}, msg: DisconnectMessage{nil}},
		{buf: []byte{0x00}, shouldFail: true},
		{buf: []byte{0x00, 0x0a}, msg: DisconnectMessage{&duration}},
	}

	for _, tt := range tests {
		var msg DisconnectMessage
		if err := msg.UnmarshalBinary(tt.buf); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else {
				if msg.Duration == nil {
					if tt.msg.Duration != nil {
						t.Errorf("Expected Duration to be %v, got nil", *tt.msg.Duration)
					}
				} else {
					if tt.msg.Duration == nil {
						t.Errorf("Expected Duration to be nil, got %v", *msg.Duration)
					} else if *msg.Duration != *tt.msg.Duration {
						t.Errorf("Expected Duration to be %v, got %v", *tt.msg.Duration, *msg.Duration)
					}
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
