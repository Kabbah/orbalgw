package mqttsn

import "testing"

func TestFlagsValue(t *testing.T) {
	tests := []struct {
		flags      Flags
		value      uint8
		shouldFail bool
	}{
		{flags: Flags{}, value: 0x00},
		{flags: Flags{TopicType: PredefinedTopicID}, value: 0x01},
		{flags: Flags{TopicType: ShortTopicName}, value: 0x02},
		{flags: Flags{TopicType: 3}, shouldFail: true},
		{flags: Flags{CleanSession: true}, value: 0x04},
		{flags: Flags{Will: true}, value: 0x08},
		{flags: Flags{Retain: true}, value: 0x10},
		{flags: Flags{QoS: -2}, shouldFail: true},
		{flags: Flags{QoS: -1}, value: 0x60},
		{flags: Flags{QoS: 2}, value: 0x40},
		{flags: Flags{QoS: 3}, shouldFail: true},
		{flags: Flags{Dup: true}, value: 0x80},
		{flags: Flags{
			TopicType:    ShortTopicName,
			CleanSession: true,
			Will:         true,
			Retain:       true,
			QoS:          -1,
			Dup:          true,
		}, value: 0xfe},
	}

	for _, tt := range tests {
		if value, err := tt.flags.Value(); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else if value != tt.value {
				t.Errorf("Expected value to be %v, got %v", tt.value, value)
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestFlagsParse(t *testing.T) {
	tests := []struct {
		value      uint8
		flags      Flags
		shouldFail bool
	}{
		{value: 0x00, flags: Flags{}},
		{value: 0x01, flags: Flags{TopicType: PredefinedTopicID}},
		{value: 0x02, flags: Flags{TopicType: ShortTopicName}},
		{value: 0x03, shouldFail: true},
		{value: 0x04, flags: Flags{CleanSession: true}},
		{value: 0x08, flags: Flags{Will: true}},
		{value: 0x10, flags: Flags{Retain: true}},
		{value: 0x20, flags: Flags{QoS: 1}},
		{value: 0x40, flags: Flags{QoS: 2}},
		{value: 0x60, flags: Flags{QoS: -1}},
		{value: 0x80, flags: Flags{Dup: true}},
		{value: 0xfe, flags: Flags{
			TopicType:    ShortTopicName,
			CleanSession: true,
			Will:         true,
			Retain:       true,
			QoS:          -1,
			Dup:          true,
		}},
		{value: 0xff, shouldFail: true},
	}

	for _, tt := range tests {
		var flags Flags
		if err := flags.Parse(tt.value); err == nil {
			if tt.shouldFail {
				t.Error("Expected error, but got nil")
			} else if flags.TopicType != tt.flags.TopicType {
				t.Errorf("Expected TopicType to be %v, got %v", tt.flags.TopicType, flags.TopicType)
			} else if flags.CleanSession != tt.flags.CleanSession {
				t.Errorf("Expected CleanSession to be %v, got %v", tt.flags.CleanSession, flags.CleanSession)
			} else if flags.Will != tt.flags.Will {
				t.Errorf("Expected Will to be %v, got %v", tt.flags.Will, flags.Will)
			} else if flags.Retain != tt.flags.Retain {
				t.Errorf("Expected Retain to be %v, got %v", tt.flags.Retain, flags.Retain)
			} else if flags.QoS != tt.flags.QoS {
				t.Errorf("Expected QoS to be %v, got %v", tt.flags.QoS, flags.QoS)
			} else if flags.Dup != tt.flags.Dup {
				t.Errorf("Expected Dup to be %v, got %v", tt.flags.Dup, flags.Dup)
			}
		} else if !tt.shouldFail {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}
