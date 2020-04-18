package main

import (
	"net"

	"github.com/google/gopacket"
	layers "github.com/google/gopacket/layers"
)

func main() {
	addr := net.UDPAddr{Port: 8090}
	u, _ := net.ListenUDP("udp", &addr)
	for {
		buff := make([]byte, 1024)
		_, clientAddr, _ := u.ReadFrom(buff)
		req := readRequestFromBuffer(buff)
		handle(u, clientAddr, req)
	}
}

func readRequestFromBuffer(buff []byte) *layers.DNS {
	packet := gopacket.NewPacket(buff, layers.LayerTypeDNS, gopacket.Default)
	dnsPacket := packet.Layer(layers.LayerTypeDNS)
	req, _ := dnsPacket.(*layers.DNS)
	return req
}

func handle(u *net.UDPConn, clientAddr net.Addr, request *layers.DNS) {
	transformRequestIntoResponse(request)
	u.WriteTo(serialize(request), clientAddr)
}

func serialize(req *layers.DNS) []byte {
	buff := gopacket.NewSerializeBuffer()
	req.SerializeTo(buff, gopacket.SerializeOptions{})
	return buff.Bytes()
}

func transformRequestIntoResponse(request *layers.DNS) *layers.DNS {
	request.AA = true
	request.Answers = createAnswers(request.Questions)
	request.ANCount = uint16(len(request.Answers))
	request.OpCode = layers.DNSOpCodeNotify
	request.QR = true
	request.ResponseCode = layers.DNSResponseCodeNoErr
	return request
}

func createAnswers(questions []layers.DNSQuestion) []layers.DNSResourceRecord {
	var answers []layers.DNSResourceRecord
	for _, q := range questions {
		answers = append(answers, createRR(q.Name))
	}
	return answers
}

func createRR(hostname []byte) layers.DNSResourceRecord {
	return layers.DNSResourceRecord{
		Type:  layers.DNSTypeA,
		IP:    findIP(hostname),
		Name:  []byte(hostname),
		Class: layers.DNSClassIN,
	}
}
