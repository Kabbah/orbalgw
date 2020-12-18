package gateway

import (
	"fmt"
	"log"
	"net"

	"github.com/Kabbah/orbalgw/internal/transport"
	"github.com/Kabbah/orbalgw/mqttsn"
)

type mqttsnAddr struct {
	Addr           net.Addr
	WirelessNodeID mqttsn.WirelessNodeID
}

// Gateway is the central structure of the MQTT-SN gateway.
type Gateway struct {
	handler MQTTSNHandler

	tp transport.DatagramTransport
}

// New creates a Gateway which uses tp as its MQTT-SN datagram transport.
//
// The supplied tp should be in its initial (or stopped) state.
func New(tp transport.DatagramTransport) *Gateway {
	gw := new(Gateway)
	gw.handler.gw = gw

	gw.tp = tp
	gw.tp.SetReceiveDatagramFunc(func(dgram transport.Datagram) {
		if err := gw.handler.Handle(dgram); err != nil {
			log.Println(fmt.Errorf("gateway: MQTT-SN handling failure: %w", err))
		}
	})

	return gw
}

// Start enables the underlying transport and starts gateway operation.
func (gw *Gateway) Start() error {
	return gw.tp.Start()
}

// Stop disables the underlying transport and stops gateway operation.
func (gw *Gateway) Stop() error {
	return gw.tp.Stop()
}

func (gw *Gateway) handleAdvertise(src *mqttsnAddr, msg *mqttsn.AdvertiseMessage) {
	log.Printf("[%v] ADVERTISE: %v", src, msg)
}

func (gw *Gateway) handleSearchGw(src *mqttsnAddr, msg *mqttsn.SearchGwMessage) {
	log.Printf("[%v] SEARCHGW: %v", src, msg)
}

func (gw *Gateway) handleGwInfo(src *mqttsnAddr, msg *mqttsn.GwInfoMessage) {
	log.Printf("[%v] GWINFO: %v", src, msg)
}

func (gw *Gateway) handleConnect(src *mqttsnAddr, msg *mqttsn.ConnectMessage) {
	log.Printf("[%v] CONNECT: %v", src, msg)
}

func (gw *Gateway) handleConnAck(src *mqttsnAddr, msg *mqttsn.ConnAckMessage) {
	log.Printf("[%v] CONNACK: %v", src, msg)
}

func (gw *Gateway) handleWillTopicReq(src *mqttsnAddr, msg *mqttsn.WillTopicReqMessage) {
	log.Printf("[%v] WILLTOPICREQ: %v", src, msg)
}

func (gw *Gateway) handleWillTopic(src *mqttsnAddr, msg *mqttsn.WillTopicMessage) {
	log.Printf("[%v] WILLTOPIC: %v", src, msg)
}

func (gw *Gateway) handleWillMsgReq(src *mqttsnAddr, msg *mqttsn.WillMsgReqMessage) {
	log.Printf("[%v] WILLMSGREQ: %v", src, msg)
}

func (gw *Gateway) handleWillMsg(src *mqttsnAddr, msg *mqttsn.WillMsgMessage) {
	log.Printf("[%v] WILLMSG: %v", src, msg)
}

func (gw *Gateway) handleRegister(src *mqttsnAddr, msg *mqttsn.RegisterMessage) {
	log.Printf("[%v] REGISTER: %v", src, msg)
}

func (gw *Gateway) handleRegAck(src *mqttsnAddr, msg *mqttsn.RegAckMessage) {
	log.Printf("[%v] REGACK: %v", src, msg)
}

func (gw *Gateway) handlePublish(src *mqttsnAddr, msg *mqttsn.PublishMessage) {
	log.Printf("[%v] PUBLISH: %v", src, msg)
}

func (gw *Gateway) handlePubAck(src *mqttsnAddr, msg *mqttsn.PubAckMessage) {
	log.Printf("[%v] PUBACK: %v", src, msg)
}

func (gw *Gateway) handlePubComp(src *mqttsnAddr, msg *mqttsn.PubCompMessage) {
	log.Printf("[%v] PUBCOMP: %v", src, msg)
}

func (gw *Gateway) handlePubRec(src *mqttsnAddr, msg *mqttsn.PubRecMessage) {
	log.Printf("[%v] PUBREC: %v", src, msg)
}

func (gw *Gateway) handlePubRel(src *mqttsnAddr, msg *mqttsn.PubRelMessage) {
	log.Printf("[%v] PUBREL: %v", src, msg)
}

func (gw *Gateway) handleSubscribe(src *mqttsnAddr, msg *mqttsn.SubscribeMessage) {
	log.Printf("[%v] SUBSCRIBE: %v", src, msg)
}

func (gw *Gateway) handleSubAck(src *mqttsnAddr, msg *mqttsn.SubAckMessage) {
	log.Printf("[%v] SUBACK: %v", src, msg)
}

func (gw *Gateway) handleUnsubscribe(src *mqttsnAddr, msg *mqttsn.UnsubscribeMessage) {
	log.Printf("[%v] UNSUBSCRIBE: %v", src, msg)
}

func (gw *Gateway) handleUnsubAck(src *mqttsnAddr, msg *mqttsn.UnsubAckMessage) {
	log.Printf("[%v] UNSUBACK: %v", src, msg)
}

func (gw *Gateway) handlePingReq(src *mqttsnAddr, msg *mqttsn.PingReqMessage) {
	log.Printf("[%v] PINGREQ: %v", src, msg)
}

func (gw *Gateway) handlePingResp(src *mqttsnAddr, msg *mqttsn.PingRespMessage) {
	log.Printf("[%v] PINGRESP: %v", src, msg)
}

func (gw *Gateway) handleDisconnect(src *mqttsnAddr, msg *mqttsn.DisconnectMessage) {
	log.Printf("[%v] DISCONNECT: %v", src, msg)
}

func (gw *Gateway) handleWillTopicUpd(src *mqttsnAddr, msg *mqttsn.WillTopicUpdMessage) {
	log.Printf("[%v] WILLTOPICUPD: %v", src, msg)
}

func (gw *Gateway) handleWillTopicResp(src *mqttsnAddr, msg *mqttsn.WillTopicRespMessage) {
	log.Printf("[%v] WILLTOPICRESP: %v", src, msg)
}

func (gw *Gateway) handleWillMsgUpd(src *mqttsnAddr, msg *mqttsn.WillMsgUpdMessage) {
	log.Printf("[%v] WILLMSGUPD: %v", src, msg)
}

func (gw *Gateway) handleWillMsgResp(src *mqttsnAddr, msg *mqttsn.WillMsgRespMessage) {
	log.Printf("[%v] WILLMSGRESP: %v", src, msg)
}
