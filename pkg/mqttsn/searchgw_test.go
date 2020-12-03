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
		buf, err := tt.msg.Marshal()
		if err == nil {
			if bytes.Compare(buf, tt.buf) != 0 {
				t.Error("Message body is wrong")
			}
		} else {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestSearchGwUnmarshal(t *testing.T) {
	tests := []struct {
		buf []byte
		msg SearchGwMessage
	}{
		{[]byte{0x0a}, SearchGwMessage{10}},
	}

	for _, tt := range tests {
		var msg SearchGwMessage
		if err := msg.Unmarshal(tt.buf); err == nil {
			if msg.Radius != tt.msg.Radius {
				t.Errorf("Expected Radius to be %v, got %v", tt.msg.Radius, msg.Radius)
			}
		} else {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
