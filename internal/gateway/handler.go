package gateway

import (
	"fmt"

	"github.com/Kabbah/orbalgw/internal/transport"
	"github.com/Kabbah/orbalgw/mqttsn"
)

// MQTTSNHandler handles datagrams received from the MQTT-SN transport.
type MQTTSNHandler struct {
}

// Handle parses a datagram into its respective MQTT-SN message and calls the
// appropriate method in the Gateway.
func (h *MQTTSNHandler) Handle(dgram transport.Datagram) error {
	var msg *mqttsn.Message

	if mqttsn.IsEncapsulated(dgram.Data) {
		var frame mqttsn.EncapsulatedMessage
		if err := frame.UnmarshalBinary(dgram.Data); err != nil {
			return err
		}
		msg = &frame.Message
	} else {
		msg = new(mqttsn.Message)
		if err := msg.UnmarshalBinary(dgram.Data); err != nil {
			return err
		}
	}

	return h.handleMessage(msg)
}

func (h *MQTTSNHandler) handleMessage(msg *mqttsn.Message) error {
	var err error
	switch msg.Type {
	case mqttsn.Advertise:
		body := new(mqttsn.AdvertiseMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.SearchGw:
		body := new(mqttsn.SearchGwMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.GwInfo:
		body := new(mqttsn.GwInfoMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.Connect:
		body := new(mqttsn.ConnectMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.ConnAck:
		body := new(mqttsn.ConnAckMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.WillTopicReq:
		body := new(mqttsn.WillTopicReqMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.WillTopic:
		body := new(mqttsn.WillTopicMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.WillMsgReq:
		body := new(mqttsn.WillMsgReqMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.WillMsg:
		body := new(mqttsn.WillMsgMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.Register:
		body := new(mqttsn.RegisterMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.RegAck:
		body := new(mqttsn.RegAckMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.Publish:
		body := new(mqttsn.PublishMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.PubAck:
		body := new(mqttsn.PubAckMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.PubComp:
		body := new(mqttsn.PubCompMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.PubRec:
		body := new(mqttsn.PubRecMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.PubRel:
		body := new(mqttsn.PubRelMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.Subscribe:
		body := new(mqttsn.SubscribeMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.SubAck:
		body := new(mqttsn.SubAckMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.Unsubscribe:
		body := new(mqttsn.UnsubscribeMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.UnsubAck:
		body := new(mqttsn.UnsubAckMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.PingReq:
		body := new(mqttsn.PingReqMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.PingResp:
		body := new(mqttsn.PingRespMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.Disconnect:
		body := new(mqttsn.DisconnectMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.WillTopicUpd:
		body := new(mqttsn.WillTopicUpdMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.WillTopicResp:
		body := new(mqttsn.WillTopicRespMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.WillMsgUpd:
		body := new(mqttsn.WillMsgUpdMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	case mqttsn.WillMsgResp:
		body := new(mqttsn.WillMsgRespMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			// TODO
		}
	default:
		err = fmt.Errorf("parser: unknown message type (%v)", msg.Type)
	}

	return err
}
