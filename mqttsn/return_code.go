package mqttsn

// The ReturnCode field in MQTT-SN is 1-byte long.
type ReturnCode uint8

// Values of the ReturnCode field.
const (
	Accepted               ReturnCode = 0x00
	RejectedCongestion     ReturnCode = 0x01
	RejectedInvalidTopicID ReturnCode = 0x02
	RejectedNotSupported   ReturnCode = 0x03
)
