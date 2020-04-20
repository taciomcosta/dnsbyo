package dns

import (
	"net"
	"testing"
)

var FindIPTests map[string]string = map[string]string{
	"google.com": "216.58.196.142",
	"amazon.com": "176.32.103.205",
}

func setup() {
	recordsFilePath = "records.json"
}

func TestFindIP(t *testing.T) {
	setup()
	for name, ip := range FindIPTests {
		got, err := FindIP(name)
		if err != nil {
			t.Error(err)
		}
		if got.String() != ip {
			t.Errorf("expected %s, got %s", ip, got)
		}
	}
}

var FindNameTests map[string]string = map[string]string{
	"216.58.196.142": "google.com",
	"176.32.103.205": "amazon.com",
}

func TestFindName(t *testing.T) {
	setup()
	for ip, name := range FindNameTests {
		got, err := FindName(net.ParseIP(ip))
		if err != nil {
			t.Error(err)
		}
		if name != got {
			t.Errorf("expected %s, got %s", name, got)
		}
	}
}
