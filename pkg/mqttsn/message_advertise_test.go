package mqttsn

import (
	"bytes"
	"testing"
)

func TestAdvertiseMarshal(t *testing.T) {
	tests := []struct {
		msg AdvertiseMessage
		buf []byte
	}{
		{AdvertiseMessage{1, 1000}, []byte{0x01, 0x03, 0xe8}},
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

func TestAdvertiseUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        AdvertiseMessage
		shouldFail bool
	}{
		{nil, AdvertiseMessage{}, true},
		{[]byte{}, AdvertiseMessage{}, true},
		{[]byte{0x01, 0x03}, AdvertiseMessage{}, true},
		{[]byte{0x01, 0x03, 0xe8}, AdvertiseMessage{1, 1000}, false},
	}

	for _, tt := range tests {
		var msg AdvertiseMessage
		if err := msg.UnmarshalBinary(tt.buf); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else {
				if msg.GatewayID != tt.msg.GatewayID {
					t.Errorf("Expected GatewayID to be %v, got %v", tt.msg.GatewayID, msg.GatewayID)
				}
				if msg.Duration != tt.msg.Duration {
					t.Errorf("Expected Duration to be %v, got %v", tt.msg.Duration, msg.Duration)
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
