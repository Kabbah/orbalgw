package mqttsn

import (
	"bytes"
	"testing"
)

func TestGwInfoMarshal(t *testing.T) {
	tests := []struct {
		msg GwInfoMessage
		buf []byte
	}{
		{GwInfoMessage{1, nil}, []byte{0x01}},
		{GwInfoMessage{1, []byte{}}, []byte{0x01}},
		{GwInfoMessage{1, []byte{0xc0, 0xa8, 0x00, 0x02}}, []byte{0x01, 0xc0, 0xa8, 0x00, 0x02}},
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

func TestGwInfoUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        GwInfoMessage
		shouldFail bool
	}{
		{buf: nil, shouldFail: true},
		{buf: []byte{}, shouldFail: true},
		{buf: []byte{0x01}, msg: GwInfoMessage{1, nil}},
		{buf: []byte{0x01, 0xc0, 0xa8, 0x00, 0x02}, msg: GwInfoMessage{1, []byte{0xc0, 0xa8, 0x00, 0x02}}},
	}

	for _, tt := range tests {
		var msg GwInfoMessage
		if err := msg.UnmarshalBinary(tt.buf); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else {
				if msg.GatewayID != tt.msg.GatewayID {
					t.Errorf("Expected GatewayID to be %v, got %v", tt.msg.GatewayID, msg.GatewayID)
				}
				if !bytes.Equal(msg.GatewayAddress, tt.msg.GatewayAddress) {
					t.Error("Message body is wrong")
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
