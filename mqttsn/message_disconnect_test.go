package mqttsn

import (
	"bytes"
	"testing"
	"time"
)

func TestDisconnectMarshal(t *testing.T) {
	duration := 10 * time.Second
	badDuration := 65536 * time.Second
	tests := []struct {
		msg        DisconnectMessage
		buf        []byte
		shouldFail bool
	}{{
		msg: DisconnectMessage{nil},
		buf: []byte{},
	}, {
		msg: DisconnectMessage{&duration},
		buf: []byte{0x00, 0x0a},
	}, {
		msg:        DisconnectMessage{&badDuration},
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

func TestDisconnectUnmarshal(t *testing.T) {
	duration := 10 * time.Second
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
