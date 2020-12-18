package gateway

import (
	"errors"
	"fmt"

	"github.com/Kabbah/orbalgw/internal/transport"
	"github.com/Kabbah/orbalgw/mqttsn"
)

// MQTTSNHandler handles datagrams received from the MQTT-SN transport.
type MQTTSNHandler struct {
	gw *Gateway
}

// Handle parses a datagram into its respective MQTT-SN message and calls the
// appropriate method in the Gateway.
func (h *MQTTSNHandler) Handle(dgram transport.Datagram) error {
	addr := mqttsnAddr{Addr: dgram.Addr}
	var msg *mqttsn.Message

	if mqttsn.IsEncapsulated(dgram.Data) {
		var frame mqttsn.EncapsulatedMessage
		if err := frame.UnmarshalBinary(dgram.Data); err != nil {
			return err
		}
		addr.WirelessNodeID = frame.WirelessNodeID
		msg = &frame.Message
	} else {
		msg = new(mqttsn.Message)
		if err := msg.UnmarshalBinary(dgram.Data); err != nil {
			return err
		}
	}

	return h.handleMessage(&addr, msg)
}

func (h *MQTTSNHandler) handleMessage(src *mqttsnAddr, msg *mqttsn.Message) error {
	gw := h.gw
	if gw == nil {
		return errors.New("handler: gateway not set")
	}

	var err error
	switch msg.Type {
	case mqttsn.Advertise:
		body := new(mqttsn.AdvertiseMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleAdvertise(src, body)
		}
	case mqttsn.SearchGw:
		body := new(mqttsn.SearchGwMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleSearchGw(src, body)
		}
	case mqttsn.GwInfo:
		body := new(mqttsn.GwInfoMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleGwInfo(src, body)
		}
	case mqttsn.Connect:
		body := new(mqttsn.ConnectMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleConnect(src, body)
		}
	case mqttsn.ConnAck:
		body := new(mqttsn.ConnAckMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleConnAck(src, body)
		}
	case mqttsn.WillTopicReq:
		body := new(mqttsn.WillTopicReqMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleWillTopicReq(src, body)
		}
	case mqttsn.WillTopic:
		body := new(mqttsn.WillTopicMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleWillTopic(src, body)
		}
	case mqttsn.WillMsgReq:
		body := new(mqttsn.WillMsgReqMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleWillMsgReq(src, body)
		}
	case mqttsn.WillMsg:
		body := new(mqttsn.WillMsgMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleWillMsg(src, body)
		}
	case mqttsn.Register:
		body := new(mqttsn.RegisterMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleRegister(src, body)
		}
	case mqttsn.RegAck:
		body := new(mqttsn.RegAckMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleRegAck(src, body)
		}
	case mqttsn.Publish:
		body := new(mqttsn.PublishMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handlePublish(src, body)
		}
	case mqttsn.PubAck:
		body := new(mqttsn.PubAckMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handlePubAck(src, body)
		}
	case mqttsn.PubComp:
		body := new(mqttsn.PubCompMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handlePubComp(src, body)
		}
	case mqttsn.PubRec:
		body := new(mqttsn.PubRecMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handlePubRec(src, body)
		}
	case mqttsn.PubRel:
		body := new(mqttsn.PubRelMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handlePubRel(src, body)
		}
	case mqttsn.Subscribe:
		body := new(mqttsn.SubscribeMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleSubscribe(src, body)
		}
	case mqttsn.SubAck:
		body := new(mqttsn.SubAckMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleSubAck(src, body)
		}
	case mqttsn.Unsubscribe:
		body := new(mqttsn.UnsubscribeMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleUnsubscribe(src, body)
		}
	case mqttsn.UnsubAck:
		body := new(mqttsn.UnsubAckMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleUnsubAck(src, body)
		}
	case mqttsn.PingReq:
		body := new(mqttsn.PingReqMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handlePingReq(src, body)
		}
	case mqttsn.PingResp:
		body := new(mqttsn.PingRespMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handlePingResp(src, body)
		}
	case mqttsn.Disconnect:
		body := new(mqttsn.DisconnectMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleDisconnect(src, body)
		}
	case mqttsn.WillTopicUpd:
		body := new(mqttsn.WillTopicUpdMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleWillTopicUpd(src, body)
		}
	case mqttsn.WillTopicResp:
		body := new(mqttsn.WillTopicRespMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleWillTopicResp(src, body)
		}
	case mqttsn.WillMsgUpd:
		body := new(mqttsn.WillMsgUpdMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleWillMsgUpd(src, body)
		}
	case mqttsn.WillMsgResp:
		body := new(mqttsn.WillMsgRespMessage)
		if err = body.UnmarshalBinary(msg.Body); err == nil {
			gw.handleWillMsgResp(src, body)
		}
	default:
		err = fmt.Errorf("handler: unknown message type (%v)", msg.Type)
	}

	return err
}
