package dns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
)

var records map[string]string = fromRecordsFile()

func fromRecordsFile() map[string]string {
	bytes, err := ioutil.ReadFile("dns/records.json")
	if err != nil {
		return map[string]string{}
	}
	records := make(map[string]string)
	json.Unmarshal(bytes, &records)
	return records
}

func findIP(name string) (net.IP, error) {
	ip, _ := records[name]
	parsedIP, _, err := net.ParseCIDR(ip + "/24")
	return parsedIP, err
}

func findName(requestedIP net.IP) (string, error) {
	for name, ip := range records {
		if requestedIP.Equal(net.ParseIP(ip)) {
			return name, nil
		}
	}
	return "", fmt.Errorf("No name found for %s", requestedIP.String())
}
