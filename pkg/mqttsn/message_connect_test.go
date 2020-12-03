package mqttsn

import (
	"bytes"
	"testing"
)

func TestConnectMarshal(t *testing.T) {
	tests := []struct {
		msg ConnectMessage
		buf []byte
	}{
		{ConnectMessage{Flags{Will: true}, 1000, "Test"}, []byte{0x08, 0x01, 0x03, 0xe8, 0x54, 0x65, 0x73, 0x74}},
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

func TestConnectUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        ConnectMessage
		shouldFail bool
	}{
		{buf: nil, shouldFail: true},
		{buf: []byte{}, shouldFail: true},
		{buf: []byte{0x00, 0x01, 0x03}, shouldFail: true},
		{buf: []byte{0x00, 0x01, 0x03, 0xe8}, msg: ConnectMessage{Flags{}, 1000, ""}},
		{buf: []byte{0x08, 0x01, 0x03, 0xe8, 0x30}, msg: ConnectMessage{Flags{Will: true}, 1000, "0"}},
		{buf: []byte{0x08, 0xff, 0x03, 0xe8, 0x30}, shouldFail: true},
	}

	for _, tt := range tests {
		var msg ConnectMessage
		if err := msg.UnmarshalBinary(tt.buf); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else {
				if msg.Flags != tt.msg.Flags {
					t.Error("Unexpected flags")
				}
				if msg.Duration != tt.msg.Duration {
					t.Errorf("Expected Duration to be %v, got %v", tt.msg.Duration, msg.Duration)
				}
				if msg.ClientID != tt.msg.ClientID {
					t.Errorf("Expected ClientID to be %v, got %v", tt.msg.ClientID, msg.ClientID)
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
