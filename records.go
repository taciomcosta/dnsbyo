package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
)

const recordsFilePath = "records.json"

func findIP(name []byte) (net.IP, error) {
	records := fromRecordsFile()
	ip, _ := records[string(name)]
	parsedIP, _, err := net.ParseCIDR(ip + "/24")
	return parsedIP, err
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

func findName(requestedIP net.IP) ([]byte, error) {
	records := fromRecordsFile()
	for name, ip := range records {
		if requestedIP.Equal(net.ParseIP(ip)) {
			return []byte(name), nil
		}
	}
	return []byte{}, fmt.Errorf("No name found for %s", requestedIP.String())
}
