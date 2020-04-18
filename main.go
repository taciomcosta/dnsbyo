package main

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	layers "github.com/google/gopacket/layers"
)

var records map[string]string = map[string]string{
	"google.com": "216.58.196.142",
	"amazon.com": "176.32.103.205",
}

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

func readRequestFromBuffer(tmp []byte) *layers.DNS {
	packet := gopacket.NewPacket(tmp, layers.LayerTypeDNS, gopacket.Default)
	dnsPacket := packet.Layer(layers.LayerTypeDNS)
	req, _ := dnsPacket.(*layers.DNS)
	return req
}

func handle(u *net.UDPConn, clientAddr net.Addr, request *layers.DNS) {
	transformRequestIntoResponse(request)
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}
	if err := request.SerializeTo(buf, opts); err != nil {
		fmt.Printf("Request serialization error")
	}
	u.WriteTo(buf.Bytes(), clientAddr)
}

func transformRequestIntoResponse(request *layers.DNS) *layers.DNS {
	request.AA = true
	request.ANCount = uint16(len(request.Questions))
	request.Answers = createAnswers(request.Questions)
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
		IP:    getIP(hostname),
		Name:  []byte(hostname),
		Class: layers.DNSClassIN,
	}
}

func getIP(requestedName []byte) net.IP {
	ip, _ := records[string(requestedName)]
	parsedIP, _, _ := net.ParseCIDR(ip + "/24")
	return parsedIP
}
