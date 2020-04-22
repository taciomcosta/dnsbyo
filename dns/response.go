package dns

import (
	"net"
)

type Response struct {
	AA      bool
	RData   ResourceRecord
	ANCount uint16
	RCode   RCode
}

type ResourceRecord struct {
	Class Class
	Type  Type
	IP    net.IP
	Name  string
}

type RCode uint16

const (
	RCodeNoErr    = 0
	RCodeNXDomain = 3
)
