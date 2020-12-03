package mqttsn

// The TopicType field in MQTT-SN is 1-byte long.
type TopicType uint8

// Values of the TopicType field.
const (
	NormalTopicID TopicType = iota
	PredefinedTopicID
	ShortTopicName
)
