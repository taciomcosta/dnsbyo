package dns

import (
	"net"
	"strings"
)

func Resolve(query Query) Response {
	response := Response{}
	response.AA = true
	response.RData = createRR(query)
	response.ANCount = ancount(response)
	response.RCode = responseCode(response)
	return response
}

func createRR(query Query) ResourceRecord {
	var rr ResourceRecord
	var err error
	if isReverseLookup(query.Name) {
		rr, err = reverseRR(query.Name)
	} else {
		rr, err = standardRR(query.Name)
	}
	if err != nil {
		return ResourceRecord{}
	}
	return rr
}

func isReverseLookup(name string) bool {
	return strings.HasSuffix(name, ".in-addr.arpa")
}

func standardRR(hostname string) (ResourceRecord, error) {
	ip, err := findIP(string(hostname))
	return existingRR(ip, hostname), err
}

func existingRR(ip net.IP, hostname string) ResourceRecord {
	return ResourceRecord{
		Type:  QTypeA,
		IP:    ip,
		Name:  hostname,
		Class: QClassIN,
	}
}

func reverseRR(ipInAddrArpa string) (ResourceRecord, error) {
	ip := parseInAddrArpa(ipInAddrArpa)
	name, err := findName(ip)
	return existingRR(ip, name), err
}

func parseInAddrArpa(ipInAddrArpa string) net.IP {
	ip := strings.Replace(ipInAddrArpa, ".in-addr.arpa", "", 1)
	octets := strings.Split(ip, ".")
	if len(octets) != 4 {
		return nil
	}
	parsedIP := net.ParseIP(octets[3] + "." + octets[2] + "." + octets[1] + "." + octets[0])
	return parsedIP
}

func responseCode(response Response) RCode {
	if response.RData.IP != nil {
		return RCodeNoErr
	}
	return RCodeNXDomain
}

func ancount(response Response) uint16 {
	if response.RData.IP != nil {
		return 1
	}
	return 0
}
