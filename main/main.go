package main

import (
	"github.com/taciomcosta/dnsbyo/dns"
	"github.com/taciomcosta/dnsbyo/packet"
	"net"
)

var conn *net.UDPConn

func main() {
	setup()
	for {
		handle()
	}
}

func setup() {
	conn, _ = net.ListenUDP("udp", &net.UDPAddr{Port: 8090})
}

func handle() {
	p := readPacket()
	response := dns.Resolve(p.Query())
	p.AddResponse(response)
	conn.WriteTo(p.Serialize(), p.ClientAddr())
}

func readPacket() packet.Packet {
	buff := make([]byte, 1024)
	_, clientAddr, _ := conn.ReadFrom(buff)
	return packet.New(buff, clientAddr)
}
