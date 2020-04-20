package dns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
)

var recordsFilePath string

func init() {
	recordsFilePath = "dns/records.json"
}

func FindIP(name string) (net.IP, error) {
	records := fromRecordsFile()
	ip, _ := records[name]
	parsedIP, _, err := net.ParseCIDR(ip + "/24")
	return parsedIP, err
}

func FindName(requestedIP net.IP) (string, error) {
	records := fromRecordsFile()
	for name, ip := range records {
		if requestedIP.Equal(net.ParseIP(ip)) {
			return name, nil
		}
	}
	return "", fmt.Errorf("No name found for %s", requestedIP.String())
}

func fromRecordsFile() map[string]string {
	bytes := readFile(recordsFilePath)
	records := make(map[string]string)
	json.Unmarshal(bytes, &records)
	return records
}

func readFile(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}
	}
	return bytes
}
