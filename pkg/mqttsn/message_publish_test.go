package mqttsn

import (
	"bytes"
	"testing"
)

func TestPublishMarshal(t *testing.T) {
	tests := []struct {
		msg PublishMessage
		buf []byte
	}{
		{PublishMessage{Flags{QoS: 1}, 1, 2, nil}, []byte{0x20, 0x00, 0x01, 0x00, 0x02}},
		{PublishMessage{Flags{QoS: 1}, 1, 2, []byte{0xab}}, []byte{0x20, 0x00, 0x01, 0x00, 0x02, 0xab}},
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

func TestPublishUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        PublishMessage
		shouldFail bool
	}{
		{buf: nil, shouldFail: true},
		{buf: []byte{}, shouldFail: true},
		{buf: []byte{0x20, 0x00, 0x01, 0x00}, shouldFail: true},
		{buf: []byte{0x20, 0x00, 0x01, 0x00, 0x02}, msg: PublishMessage{Flags{QoS: 1}, 1, 2, nil}},
		{buf: []byte{0x20, 0x00, 0x01, 0x00, 0x02, 0xab}, msg: PublishMessage{Flags{QoS: 1}, 1, 2, []byte{0xab}}},
	}

	for _, tt := range tests {
		var msg PublishMessage
		if err := msg.UnmarshalBinary(tt.buf); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else {
				if msg.Flags != tt.msg.Flags {
					t.Error("Unexpected flags")
				}
				if msg.TopicID != tt.msg.TopicID {
					t.Errorf("Expected TopicID to be %v, got %v", tt.msg.TopicID, msg.TopicID)
				}
				if msg.MessageID != tt.msg.MessageID {
					t.Errorf("Expected MessageID to be %v, got %v", tt.msg.MessageID, msg.MessageID)
				}
				if !bytes.Equal(msg.Data, tt.msg.Data) {
					t.Error("Data is wrong")
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
