package mqttsn

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
