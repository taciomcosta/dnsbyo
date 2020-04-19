package packet

import (
	"github.com/google/gopacket"
	layers "github.com/google/gopacket/layers"
)

func NewDNSPacket(buff []byte) *layers.DNS {
	packet := gopacket.NewPacket(buff, layers.LayerTypeDNS, gopacket.Default)
	layer := packet.Layer(layers.LayerTypeDNS)
	dnsPacket, _ := layer.(*layers.DNS)
	return dnsPacket
}

func Serialize(req *layers.DNS) []byte {
	buff := gopacket.NewSerializeBuffer()
	req.SerializeTo(buff, gopacket.SerializeOptions{})
	return buff.Bytes()
}