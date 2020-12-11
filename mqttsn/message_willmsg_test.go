package mqttsn

import (
	"bytes"
	"testing"
)

func TestWillMsgMarshal(t *testing.T) {
	tests := []struct {
		msg WillMsgMessage
		buf []byte
	}{
		{WillMsgMessage{}, []byte{}},
		{WillMsgMessage{[]byte{0x01, 0x02, 0x03}}, []byte{0x01, 0x02, 0x03}},
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

func TestWillMsgUnmarshal(t *testing.T) {
	tests := []struct {
		buf []byte
		msg WillMsgMessage
	}{
		{[]byte{}, WillMsgMessage{}},
		{[]byte{0x01, 0x02, 0x03}, WillMsgMessage{[]byte{0x01, 0x02, 0x03}}},
	}

	for _, tt := range tests {
		var msg WillMsgMessage
		err := msg.UnmarshalBinary(tt.buf)
		if err == nil {
			if !bytes.Equal(msg.Data, tt.msg.Data) {
				t.Error("Message data is wrong")
			}
		} else {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
