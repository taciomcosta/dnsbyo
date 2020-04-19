package main

import (
	"net"

	"github.com/google/gopacket"
	layers "github.com/google/gopacket/layers"
)

var conn *net.UDPConn

func main() {
	setup()
	for {
		buff := make([]byte, 1024)
		_, clientAddr, _ := conn.ReadFrom(buff)
		req := readRequestFromBuffer(buff)
		handle(clientAddr, req)
	}
}

func setup() {
	conn, _ = net.ListenUDP("udp", &net.UDPAddr{Port: 8090})
}

func readRequestFromBuffer(buff []byte) *layers.DNS {
	packet := gopacket.NewPacket(buff, layers.LayerTypeDNS, gopacket.Default)
	dnsPacket := packet.Layer(layers.LayerTypeDNS)
	req, _ := dnsPacket.(*layers.DNS)
	return req
}

func handle(clientAddr net.Addr, query *layers.DNS) {
	transformQueryIntoResponse(query)
	conn.WriteTo(serialize(query), clientAddr)
}

func serialize(req *layers.DNS) []byte {
	buff := gopacket.NewSerializeBuffer()
	req.SerializeTo(buff, gopacket.SerializeOptions{})
	return buff.Bytes()
}
