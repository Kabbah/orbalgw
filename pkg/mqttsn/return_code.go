package mqttsn

// The ReturnCode field in MQTT-SN is 1-byte long.
type ReturnCode uint8

// Values of the ReturnCode field.
const (
	Accepted ReturnCode = iota
	RejectedCongestion
	RejectedInvalidTopicID
	RejectedNotSupported
)
