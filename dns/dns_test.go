package dns

import (
	"testing"
)

var testsResolve map[string]string = map[string]string{
	"google.com": "216.58.196.142",
	"amazon.com": "176.32.103.205",
}

func TestResolve(t *testing.T) {
	records = testsResolve
	for name, ip := range testsResolve {
		r := Resolve(Query{Name: name})
		if r.RData.IP.String() != ip {
			t.Errorf("expected %s, got %s", ip, r.RData.IP.String())
		}
	}
}

var testsResolveReverse map[string]string = map[string]string{
	"142.196.58.216.in-addr.arpa": "google.com",
	"205.103.32.176.in-addr.arpa": "amazon.com",
}

func TestResolveReverse(t *testing.T) {
	records = testsResolve
	for ip, name := range testsResolveReverse {
		r := Resolve(Query{Name: ip})
		if r.RData.Name != name {
			t.Errorf("expected %s, got %s", name, r.RData.Name)
		}
	}
}
