package main

import (
	"github.com/taciomcosta/dnsbyo/packet"
	"net"
)

var conn *net.UDPConn

func main() {
	setup()
	for {
		handleRequest()
	}
}

func setup() {
	conn, _ = net.ListenUDP("udp", &net.UDPAddr{Port: 8090})
}

func handleRequest() {
	buff := make([]byte, 1024)
	_, clientAddr, _ := conn.ReadFrom(buff)
	dnsPacket := packet.NewDNSPacket(buff)

	packet.TransformQueryIntoResponse(dnsPacket)
	conn.WriteTo(packet.Serialize(dnsPacket), clientAddr)
}
