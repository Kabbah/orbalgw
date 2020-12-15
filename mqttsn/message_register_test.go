package mqttsn

import (
	"bytes"
	"testing"
)

func TestRegisterMarshal(t *testing.T) {
	tests := []struct {
		msg RegisterMessage
		buf []byte
	}{
		{RegisterMessage{1, 2, "t"}, []byte{0x00, 0x01, 0x00, 0x02, 0x74}},
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

func TestRegisterUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        RegisterMessage
		shouldFail bool
	}{
		{buf: nil, shouldFail: true},
		{buf: []byte{}, shouldFail: true},
		{buf: []byte{0x00, 0x01, 0x00}, shouldFail: true},
		{buf: []byte{0x00, 0x01, 0x00, 0x02}, msg: RegisterMessage{1, 2, ""}},
		{buf: []byte{0x00, 0x01, 0x00, 0x02, 0x30}, msg: RegisterMessage{1, 2, "0"}},
	}

	for _, tt := range tests {
		var msg RegisterMessage
		if err := msg.UnmarshalBinary(tt.buf); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else {
				if msg.TopicID != tt.msg.TopicID {
					t.Errorf("Expected TopicID to be %v, got %v", tt.msg.TopicID, msg.TopicID)
				}
				if msg.MessageID != tt.msg.MessageID {
					t.Errorf("Expected MessageID to be %v, got %v", tt.msg.MessageID, msg.MessageID)
				}
				if msg.TopicName != tt.msg.TopicName {
					t.Errorf("Expected TopicName to be %v, got %v", tt.msg.TopicName, msg.TopicName)
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
