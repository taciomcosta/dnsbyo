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
		tmp := make([]byte, 1024)
		_, addr, _ := u.ReadFrom(tmp)
		clientAddr := addr
		packet := gopacket.NewPacket(tmp, layers.LayerTypeDNS, gopacket.Default)
		dnsPacket := packet.Layer(layers.LayerTypeDNS)
		tcp, _ := dnsPacket.(*layers.DNS)
		serveDNS(u, clientAddr, tcp)
	}
}

func serveDNS(u *net.UDPConn, clientAddr net.Addr, request *layers.DNS) {
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}
	addReplyMessageToRequest(request)
	if err := request.SerializeTo(buf, opts); err != nil {
		panic(err)
	}
	u.WriteTo(buf.Bytes(), clientAddr)
}

func addReplyMessageToRequest(request *layers.DNS) *layers.DNS {
	request.QR = true
	request.ANCount = 1
	request.OpCode = layers.DNSOpCodeNotify
	request.AA = true
	request.Answers = append(request.Answers, createRR(request.Questions[0].Name))
	request.ResponseCode = layers.DNSResponseCodeNoErr
	return request
}

func createRR(requestedName []byte) layers.DNSResourceRecord {
	return layers.DNSResourceRecord{
		Type:  layers.DNSTypeA,
		IP:    getIP(requestedName),
		Name:  []byte(requestedName),
		Class: layers.DNSClassIN,
	}
}

func getIP(requestedName []byte) net.IP {
	ip, ok := records[string(requestedName)]
	if !ok {
		fmt.Printf("%s -> NOT FOUND\n", requestedName)
	} else {
		fmt.Printf("%s -> %s\n", requestedName, ip)
	}
	parsedIP, _, _ := net.ParseCIDR(ip + "/24")
	return parsedIP
}
