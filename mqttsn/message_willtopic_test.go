package mqttsn

import (
	"bytes"
	"testing"
)

func TestWillTopicMarshal(t *testing.T) {
	tests := []struct {
		msg WillTopicMessage
		buf []byte
	}{
		{WillTopicMessage{Empty: true}, []byte{}},
		{WillTopicMessage{Flags: Flags{}, TopicName: "dead"}, []byte{0x00, 0x64, 0x65, 0x61, 0x64}},
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

func TestWillTopicUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        WillTopicMessage
		shouldFail bool
	}{
		{buf: nil, msg: WillTopicMessage{Empty: true}},
		{buf: []byte{}, msg: WillTopicMessage{Empty: true}},
		{buf: []byte{0x00}, shouldFail: true},
		{buf: []byte{0x00, 0x64, 0x65, 0x61, 0x64}, msg: WillTopicMessage{Flags: Flags{}, TopicName: "dead"}},
	}

	for _, tt := range tests {
		var msg WillTopicMessage
		if err := msg.UnmarshalBinary(tt.buf); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else {
				if msg.Empty != tt.msg.Empty {
					t.Errorf("Expected Empty to be %v, but got %v", tt.msg.Empty, msg.Empty)
				}
				if msg.Flags != tt.msg.Flags {
					t.Error("Unexpected flags")
				}
				if msg.TopicName != tt.msg.TopicName {
					t.Errorf("Expected TopicName to be %v, but got %v", tt.msg.TopicName, msg.TopicName)
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
