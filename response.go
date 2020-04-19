package main

import (
	"net"
	"strings"

	layers "github.com/google/gopacket/layers"
)

func transformQueryIntoResponse(query *layers.DNS) *layers.DNS {
	query.AA = true
	query.Answers = createAnswers(query)
	query.ANCount = uint16(len(query.Answers))
	query.ResponseCode = responseCode(query)
	query.QR = true
	return query
}

func createAnswers(req *layers.DNS) []layers.DNSResourceRecord {
	var answers []layers.DNSResourceRecord
	for i, _ := range req.Questions {
		if rr, err := createRR(req, i); err == nil {
			answers = append(answers, rr)
		}
	}
	return answers
}

func createRR(query *layers.DNS, i int) (layers.DNSResourceRecord, error) {
	if isReverseLookup(query.Questions[0].Name) {
		return reverseRR(query.Questions[0].Name)
	}
	return standardRR(query.Questions[0].Name)
}

func isReverseLookup(name []byte) bool {
	return strings.HasSuffix(string(name), ".in-addr.arpa")
}

func standardRR(hostname []byte) (layers.DNSResourceRecord, error) {
	ip, err := findIP(string(hostname))
	if err != nil {
		return layers.DNSResourceRecord{}, err
	}
	return existingRR(ip, hostname), nil
}

func existingRR(ip net.IP, hostname []byte) layers.DNSResourceRecord {
	return layers.DNSResourceRecord{
		Type:  layers.DNSTypeA,
		IP:    ip,
		Name:  []byte(hostname),
		Class: layers.DNSClassIN,
	}
}

func responseCode(query *layers.DNS) layers.DNSResponseCode {
	if len(query.Answers) > 0 {
		return layers.DNSResponseCodeNoErr
	}
	return layers.DNSResponseCodeNXDomain
}

func reverseRR(ipInAddrArpa []byte) (layers.DNSResourceRecord, error) {
	ip := parseInAddrArpa(ipInAddrArpa)
	name, err := findName(ip)
	if err != nil {
		return layers.DNSResourceRecord{}, err
	}
	return existingRR(ip, name), nil
}

func parseInAddrArpa(ipInAddrArpa []byte) net.IP {
	ip := strings.Replace(string(ipInAddrArpa), ".in-addr.arpa", "", 1)
	octets := strings.Split(ip, ".")
	if len(octets) != 4 {
		return nil
	}
	parsedIP := net.ParseIP(octets[3] + "." + octets[2] + "." + octets[1] + "." + octets[0])
	return parsedIP
}
