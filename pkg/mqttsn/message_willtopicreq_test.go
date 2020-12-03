package mqttsn

import (
	"bytes"
	"testing"
)

func TestWillTopicReqMarshal(t *testing.T) {
	tests := []struct {
		msg WillTopicReqMessage
		buf []byte
	}{
		{WillTopicReqMessage{}, []byte{}},
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

func TestWillTopicReqUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        WillTopicReqMessage
		shouldFail bool
	}{
		{buf: nil, msg: WillTopicReqMessage{}},
		{buf: []byte{}, msg: WillTopicReqMessage{}},
		{buf: []byte{0x00}, shouldFail: true},
	}

	for _, tt := range tests {
		var msg WillTopicReqMessage
		err := msg.UnmarshalBinary(tt.buf)
		if err == nil && tt.shouldFail {
			t.Error("Expected error, but got nil")
		} else if err != nil && !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
