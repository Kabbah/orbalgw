package mqttsn

import (
	"bytes"
	"testing"
)

func TestWillMsgReqMarshal(t *testing.T) {
	tests := []struct {
		msg WillMsgReqMessage
		buf []byte
	}{
		{WillMsgReqMessage{}, []byte{}},
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

func TestWillMsgReqUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        WillMsgReqMessage
		shouldFail bool
	}{
		{buf: nil, msg: WillMsgReqMessage{}},
		{buf: []byte{}, msg: WillMsgReqMessage{}},
		{buf: []byte{0x00}, shouldFail: true},
	}

	for _, tt := range tests {
		var msg WillMsgReqMessage
		err := msg.UnmarshalBinary(tt.buf)
		if err == nil && tt.shouldFail {
			t.Error("Expected error, but got nil")
		} else if err != nil && !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
