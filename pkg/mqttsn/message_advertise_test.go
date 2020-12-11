package mqttsn

import (
	"bytes"
	"testing"
	"time"
)

func TestAdvertiseMarshal(t *testing.T) {
	tests := []struct {
		msg        AdvertiseMessage
		buf        []byte
		shouldFail bool
	}{{
		msg: AdvertiseMessage{1, 1000 * time.Second},
		buf: []byte{0x01, 0x03, 0xe8},
	}, {
		msg:        AdvertiseMessage{1, 65536 * time.Second},
		shouldFail: true,
	}}

	for _, tt := range tests {
		if buf, err := tt.msg.MarshalBinary(); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else if !bytes.Equal(buf, tt.buf) {
				t.Error("Message body is wrong")
			}
		} else if !tt.shouldFail {
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
		{buf: nil, shouldFail: true},
		{buf: []byte{}, shouldFail: true},
		{buf: []byte{0x01, 0x03}, shouldFail: true},
		{buf: []byte{0x01, 0x03, 0xe8}, msg: AdvertiseMessage{1, 1000 * time.Second}},
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
