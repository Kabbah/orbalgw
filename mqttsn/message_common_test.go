package mqttsn

import (
	"bytes"
	"testing"
)

func TestEmptyMarshal(t *testing.T) {
	tests := []struct {
		msg emptyMessage
		buf []byte
	}{
		{emptyMessage{}, []byte{}},
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

func TestEmptyUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        emptyMessage
		shouldFail bool
	}{
		{buf: nil, msg: emptyMessage{}},
		{buf: []byte{}, msg: emptyMessage{}},
		{buf: []byte{0x00}, shouldFail: true},
	}

	for _, tt := range tests {
		var msg emptyMessage
		err := msg.UnmarshalBinary(tt.buf)
		if err == nil && tt.shouldFail {
			t.Error("Expected error, but got nil")
		} else if err != nil && !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestMessageIDOnlyMarshal(t *testing.T) {
	tests := []struct {
		msg messageIDOnlyMessage
		buf []byte
	}{
		{messageIDOnlyMessage{1}, []byte{0x00, 0x01}},
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

func TestMessageIDOnlyUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        messageIDOnlyMessage
		shouldFail bool
	}{
		{buf: nil, shouldFail: true},
		{buf: []byte{}, shouldFail: true},
		{buf: []byte{0x00}, shouldFail: true},
		{buf: []byte{0x00, 0x01}, msg: messageIDOnlyMessage{1}},
	}

	for _, tt := range tests {
		var msg messageIDOnlyMessage
		if err := msg.UnmarshalBinary(tt.buf); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else {
				if msg.MessageID != tt.msg.MessageID {
					t.Errorf("Expected MessageID to be %v, got %v", tt.msg.MessageID, msg.MessageID)
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestReturnCodeOnlyMarshal(t *testing.T) {
	tests := []struct {
		msg returnCodeOnlyMessage
		buf []byte
	}{
		{returnCodeOnlyMessage{Accepted}, []byte{0x00}},
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

func TestReturnCodeOnlyUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        returnCodeOnlyMessage
		shouldFail bool
	}{
		{buf: nil, shouldFail: true},
		{buf: []byte{}, shouldFail: true},
		{buf: []byte{0x00}, msg: returnCodeOnlyMessage{Accepted}},
		{buf: []byte{0x00, 0x00}, shouldFail: true},
	}

	for _, tt := range tests {
		var msg returnCodeOnlyMessage
		if err := msg.UnmarshalBinary(tt.buf); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else {
				if msg.ReturnCode != tt.msg.ReturnCode {
					t.Errorf("Expected ReturnCode to be %v, got %v", tt.msg.ReturnCode, msg.ReturnCode)
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestTopicAckMarshal(t *testing.T) {
	tests := []struct {
		msg topicAckMessage
		buf []byte
	}{
		{topicAckMessage{1, 2, Accepted}, []byte{0x00, 0x01, 0x00, 0x02, 0x00}},
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

func TestTopicAckUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        topicAckMessage
		shouldFail bool
	}{
		{buf: nil, shouldFail: true},
		{buf: []byte{}, shouldFail: true},
		{buf: []byte{0x00, 0x01, 0x00, 0x02}, shouldFail: true},
		{buf: []byte{0x00, 0x01, 0x00, 0x02, 0x00}, msg: topicAckMessage{1, 2, Accepted}},
	}

	for _, tt := range tests {
		var msg topicAckMessage
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
				if msg.ReturnCode != tt.msg.ReturnCode {
					t.Errorf("Expected ReturnCode to be %v, got %v", tt.msg.ReturnCode, msg.ReturnCode)
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
