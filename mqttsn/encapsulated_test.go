package mqttsn

import (
	"bytes"
	"testing"
)

func TestEncapsulatedMarshal(t *testing.T) {
	tests := []struct {
		msg        EncapsulatedMessage
		buf        []byte
		shouldFail bool
	}{{
		msg: EncapsulatedMessage{
			Radius:         1,
			WirelessNodeID: []byte{0x12, 0x34, 0x56, 0x78},
			Message:        Message{PingResp, nil},
		},
		buf: []byte{0x07, 0xfe, 0x01, 0x12, 0x34, 0x56, 0x78, 0x02, 0x17},
	}, {
		msg: EncapsulatedMessage{
			Radius:         4,
			WirelessNodeID: []byte{0x12, 0x34, 0x56, 0x78},
			Message:        Message{PingResp, nil},
		},
		shouldFail: true,
	}, {
		msg: EncapsulatedMessage{
			Radius:         1,
			WirelessNodeID: make([]byte, 253),
			Message:        Message{PingResp, nil},
		},
		shouldFail: true,
	}}

	for _, tt := range tests {
		if buf, err := tt.msg.MarshalBinary(); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else if !bytes.Equal(buf, tt.buf) {
				t.Error("Frame body is wrong")
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestEncapsulatedUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        EncapsulatedMessage
		shouldFail bool
	}{{
		buf: nil, shouldFail: true,
	}, {
		buf: []byte{}, shouldFail: true,
	}, {
		buf: []byte{0x02, 0xfe}, shouldFail: true,
	}, {
		buf: []byte{0x03, 0xfe, 0x01}, shouldFail: true,
	}, {
		buf: []byte{0x03, 0xfe, 0x01, 0x02, 0x17},
		msg: EncapsulatedMessage{
			Radius:         1,
			WirelessNodeID: nil,
			Message:        Message{PingResp, nil},
		},
	}, {
		buf: []byte{0x03, 0xff, 0x01, 0x02, 0x17}, shouldFail: true,
	}, {
		buf: []byte{0xff, 0xfe, 0x01, 0x02, 0x17}, shouldFail: true,
	}, {
		buf: []byte{0x04, 0xfe, 0x01, 0xab, 0x02, 0x17},
		msg: EncapsulatedMessage{
			Radius:         1,
			WirelessNodeID: []byte{0xab},
			Message:        Message{PingResp, nil},
		},
	}}

	for _, tt := range tests {
		var msg EncapsulatedMessage
		if err := msg.UnmarshalBinary(tt.buf); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else {
				if msg.Radius != tt.msg.Radius {
					t.Errorf("Expected Radius to be %v, got %v", tt.msg.Radius, msg.Radius)
				}
				if !bytes.Equal(msg.WirelessNodeID, tt.msg.WirelessNodeID) {
					t.Error("WirelessNodeID is wrong")
				}
				if msg.Type != tt.msg.Type {
					t.Errorf("Expected Type to be %v, got %v", tt.msg.Type, msg.Type)
				}
				if !bytes.Equal(msg.Body, tt.msg.Body) {
					t.Error("Message body is wrong")
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
