package mqttsn

import (
	"bytes"
	"testing"
)

func TestSearchGwMarshal(t *testing.T) {
	tests := []struct {
		msg SearchGwMessage
		buf []byte
	}{
		{SearchGwMessage{10}, []byte{0x0a}},
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

func TestSearchGwUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        SearchGwMessage
		shouldFail bool
	}{
		{buf: nil, shouldFail: true},
		{buf: []byte{}, shouldFail: true},
		{buf: []byte{0x0a}, msg: SearchGwMessage{10}},
	}

	for _, tt := range tests {
		var msg SearchGwMessage
		if err := msg.UnmarshalBinary(tt.buf); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else {
				if msg.Radius != tt.msg.Radius {
					t.Errorf("Expected Radius to be %v, got %v", tt.msg.Radius, msg.Radius)
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
