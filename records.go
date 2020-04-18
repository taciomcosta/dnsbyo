package main

import (
	"encoding/json"
	"io/ioutil"
	"net"
)

func findIP(name []byte) (net.IP, error) {
	records := fromJSON("records.json")
	ip, _ := records[string(name)]
	parsedIP, _, err := net.ParseCIDR(ip + "/24")
	return parsedIP, err
}

func fromJSON(path string) map[string]string {
	bytes := readFile(path)
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
