package mqttsn

import (
	"bytes"
	"testing"
)

func TestSubscribeMarshal(t *testing.T) {
	tests := []struct {
		msg        SubscribeMessage
		buf        []byte
		shouldFail bool
	}{{
		msg: SubscribeMessage{MessageID: 2},
		buf: []byte{0x00, 0x00, 0x02},
	}, {
		msg: SubscribeMessage{
			MessageID: 2,
			TopicName: "012",
		},
		buf: []byte{0x00, 0x00, 0x02, 0x30, 0x31, 0x32},
	}, {
		msg: SubscribeMessage{
			Flags:     Flags{TopicType: PredefinedTopicID},
			TopicID:   1,
			MessageID: 2,
		},
		buf: []byte{0x01, 0x00, 0x02, 0x00, 0x01},
	}, {
		msg: SubscribeMessage{
			Flags:     Flags{TopicType: ShortTopicName},
			MessageID: 2,
			TopicName: "01",
		},
		buf: []byte{0x02, 0x00, 0x02, 0x30, 0x31},
	}, {
		msg: SubscribeMessage{
			Flags:     Flags{TopicType: ShortTopicName},
			MessageID: 2,
			TopicName: "012",
		},
		shouldFail: true,
	}}

	for _, tt := range tests {
		if buf, err := tt.msg.MarshalBinary(); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else if !bytes.Equal(buf, tt.buf) {
				t.Error("Message body is wrong")
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestSubscribeUnmarshal(t *testing.T) {
	tests := []struct {
		buf        []byte
		msg        SubscribeMessage
		shouldFail bool
	}{{
		buf: nil, shouldFail: true,
	}, {
		buf: []byte{}, shouldFail: true,
	}, {
		buf: []byte{}, shouldFail: true,
	}, {
		buf: []byte{0x00, 0x00, 0x02},
		msg: SubscribeMessage{MessageID: 2},
	}, {
		buf: []byte{0x00, 0x00, 0x02, 0x30, 0x31, 0x32},
		msg: SubscribeMessage{
			MessageID: 2,
			TopicName: "012",
		},
	}, {
		buf: []byte{0x01, 0x00, 0x02, 0x00, 0x01},
		msg: SubscribeMessage{
			Flags:     Flags{TopicType: PredefinedTopicID},
			MessageID: 2,
			TopicID:   1,
		},
	}, {
		buf: []byte{0x01, 0x00, 0x02, 0x30, 0x31, 0x32}, shouldFail: true,
	}, {
		buf: []byte{0x02, 0x00, 0x02, 0x30, 0x31},
		msg: SubscribeMessage{
			Flags:     Flags{TopicType: ShortTopicName},
			MessageID: 2,
			TopicName: "01",
		},
	}, {
		buf: []byte{0x02, 0x00, 0x02, 0x30, 0x31, 0x32}, shouldFail: true,
	}, {
		buf: []byte{0x03, 0x00, 0x02, 0x30, 0x31}, shouldFail: true,
	}}

	for _, tt := range tests {
		var msg SubscribeMessage
		if err := msg.UnmarshalBinary(tt.buf); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else {
				if msg.Flags != tt.msg.Flags {
					t.Error("Unexpected flags")
				} else if msg.Flags.TopicType == PredefinedTopicID {
					if msg.TopicID != tt.msg.TopicID {
						t.Errorf("Expected TopicID to be %v, got %v", tt.msg.TopicID, msg.TopicID)
					}
				} else {
					if msg.TopicName != tt.msg.TopicName {
						t.Errorf("Expected TopicName to be %v, got %v", tt.msg.TopicName, msg.TopicName)
					}
				}
				if msg.MessageID != tt.msg.MessageID {
					t.Errorf("Expected MessageID to be %v, got %v", tt.msg.MessageID, msg.MessageID)
				}
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
