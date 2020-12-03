package mqttsn

// The MessageType field in MQTT-SN is 1-byte long.
type MessageType uint8

// Values of the MessageType field.
const (
	Advertise MessageType = iota
	SearchGw
	GwInfo
	_
	Connect
	ConnAck
	WillTopicReq
	WillTopic
	WillMsgReq
	WillMsg
	Register
	RegAck
	Publish
	PubAck
	PubComp
	PubRec
	PubRel
	_
	Subscribe
	SubAck
	Unsubscribe
	UnsubAck
	PingReq
	PingResp
	Disconnect
	_
	WillTopicUpd
	WillTopicResp
	WillMsgUpd
	WillMsgResp
	Encapsulated MessageType = 254
)
