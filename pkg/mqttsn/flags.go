package mqttsn

import "fmt"

// Flags used in some MQTT-SN messages.
type Flags struct {
	Dup          bool
	QoS          int8
	Retain       bool
	Will         bool
	CleanSession bool
	TopicType    TopicType
}

// Value returns the encoded form of MQTT-SN flags.
func (f *Flags) Value() (uint8, error) {
	if f.QoS < -1 || f.QoS > 2 {
		return 0, fmt.Errorf("mqttsn: invalid QoS level (%v)", f.QoS)
	}

	bits := uint8(f.TopicType)
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

// Parse interprets the encoded form of MQTT-SN flags.
func (f *Flags) Parse(value uint8) error {
	topicType := TopicType(value & 0b0000_0011)
	if topicType > ShortTopicName {
		return fmt.Errorf("mqttsn: invalid topic type (%v)", topicType)
	}

	qos := int8((value & 0b0110_0000) >> 5)
	if qos == 3 {
		qos = -1
	}

	f.TopicType = topicType
	f.CleanSession = (value & 0b0000_0100) != 0
	f.Will = (value & 0b0000_1000) != 0
	f.Retain = (value & 0b0001_0000) != 0
	f.QoS = qos
	f.Dup = (value & 0b1000_0000) != 0
	return nil
}
