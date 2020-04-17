package main

import (
	"fmt"
	"github.com/google/gopacket"
	layers "github.com/google/gopacket/layers"
	"net"
)

var records map[string]string

func main() {
	records = map[string]string{
		"google.com": "216.58.196.142",
		"amazon.com": "176.32.103.205",
	}

	// Listen on UDP Port
	addr := net.UDPAddr{
		Port: 8090,
		IP:   net.ParseIP("127.0.0.1"),
	}
	u, _ := net.ListenUDP("udp", &addr)

	// Wait to get request on that port
	for {
		fmt.Println("waiting...")
		tmp := make([]byte, 1024)
		_, addr, _ := u.ReadFrom(tmp)
		clientAddr := addr
		fmt.Println(clientAddr)
		packet := gopacket.NewPacket(tmp, layers.LayerTypeDNS, gopacket.Default)
		dnsPacket := packet.Layer(layers.LayerTypeDNS)
		tcp, _ := dnsPacket.(*layers.DNS)
		serveDNS(u, clientAddr, tcp)
	}
}

func serveDNS(u *net.UDPConn, clientAddr net.Addr, request *layers.DNS) {
	replyMess := request
	var dnsAnswer layers.DNSResourceRecord
	dnsAnswer.Type = layers.DNSTypeA
	ip, ok := records[string(request.Questions[0].Name)]
	if !ok {
		fmt.Printf("No data present for %v\n", request.Questions[0].Name)
	}
	a, _, _ := net.ParseCIDR(ip + "/24")

	dnsAnswer.Type = layers.DNSTypeA
	dnsAnswer.IP = a
	dnsAnswer.Name = []byte(request.Questions[0].Name)
	dnsAnswer.Class = layers.DNSClassIN
	fmt.Println(request.Questions[0].Name)

	replyMess.QR = true
	replyMess.ANCount = 1
	replyMess.OpCode = layers.DNSOpCodeNotify
	replyMess.AA = true
	replyMess.Answers = append(replyMess.Answers, dnsAnswer)
	replyMess.ResponseCode = layers.DNSResponseCodeNoErr

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}
	if err := replyMess.SerializeTo(buf, opts); err != nil {
		panic(err)
	}
	u.WriteTo(buf.Bytes(), clientAddr)
}
