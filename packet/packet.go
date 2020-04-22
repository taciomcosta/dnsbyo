package packet

import (
	"github.com/google/gopacket"
	layers "github.com/google/gopacket/layers"
	"github.com/taciomcosta/dnsbyo/dns"
	"net"
)

type Packet struct {
	dnsPacket  *layers.DNS
	clientAddr net.Addr
}

func New(buff []byte, addr net.Addr) Packet {
	packet := gopacket.NewPacket(buff, layers.LayerTypeDNS, gopacket.Default)
	layer := packet.Layer(layers.LayerTypeDNS)
	dnsPacket, _ := layer.(*layers.DNS)
	return Packet{dnsPacket, addr}
}

func (p *Packet) Query() dns.Query {
	return dns.Query{
		Name:   string(p.dnsPacket.Questions[0].Name),
		QClass: dns.Class(p.dnsPacket.Questions[0].Class),
		QType:  dns.Type(p.dnsPacket.Questions[0].Type),
	}
}

func (p *Packet) Serialize() []byte {
	buff := gopacket.NewSerializeBuffer()
	p.dnsPacket.SerializeTo(buff, gopacket.SerializeOptions{})
	return buff.Bytes()
}

func (p *Packet) ClientAddr() net.Addr {
	return p.clientAddr
}

func (p *Packet) AddResponse(response dns.Response) {
	p.dnsPacket.AA = response.AA
	p.dnsPacket.ANCount = response.ANCount
	p.dnsPacket.ResponseCode = layers.DNSResponseCode(response.RCode)
	p.dnsPacket.QR = true
	p.dnsPacket.Answers = []layers.DNSResourceRecord{
		{
			Name:  []byte(response.RData.Name),
			IP:    response.RData.IP,
			Type:  layers.DNSType(response.RData.Type),
			Class: layers.DNSClass(response.RData.Class),
		},
	}
}
