package mqttsn

// The TopicType field in MQTT-SN is 1-byte long.
type TopicType uint8

// Values of the TopicType field.
const (
	NormalTopicID     TopicType = 0x00
	NormalTopicName   TopicType = 0x00
	PredefinedTopicID TopicType = 0x01
	ShortTopicName    TopicType = 0x02
)
