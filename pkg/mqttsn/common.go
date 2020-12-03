package mqttsn

import "fmt"

// ProtocolID of MQTT-SN is always 0x01.
const ProtocolID uint8 = 1

// The TopicIDType field in MQTT-SN is 1-byte long.
type TopicIDType uint8

// Values of the TopicIDType field.
const (
	NormalTopicID TopicIDType = iota
	PredefinedTopicID
	ShortTopicName
)

// The ReturnCode field in MQTT-SN is 1-byte long.
type ReturnCode uint8

// Values of the ReturnCode field.
const (
	Accepted ReturnCode = iota
	RejectedCongestion
	RejectedInvalidTopicID
	RejectedNotSupported
)

// Flags used in some MQTT-SN messages.
type Flags struct {
	Dup          bool
	QoS          int8
	Retain       bool
	Will         bool
	CleanSession bool
	TopicID      TopicIDType
}

// Value returns the encoded form of MQTT-SN flags.
func (f *Flags) Value() (uint8, error) {
	if f.QoS < -1 || f.QoS > 2 {
		return 0, fmt.Errorf("mqttsn: invalid QoS level (%v)", f.QoS)
	}

	bits := uint8(f.TopicID)
	if f.CleanSession {
		bits |= 1 << 2
	}
	if f.Will {
		bits |= 1 << 3
	}
	if f.Retain {
		bits |= 1 << 4
	}
	if f.QoS == -1 {
		bits |= 3 << 5
	} else {
		bits |= uint8(f.QoS) << 5
	}
	if f.Dup {
		bits |= 1 << 7
	}
	return bits, nil
}
