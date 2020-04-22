package dns

import (
	"net"
	"os"
	"testing"
)

var FindIPTests map[string]string = map[string]string{
	"google.com": "216.58.196.142",
	"amazon.com": "176.32.103.205",
}

func TestMain(m *testing.M) {
	recordsFilePath = "records.json"
	os.Exit(m.Run())
}

func TestFindIP(t *testing.T) {
	for name, ip := range FindIPTests {
		got, err := findIP(name)
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
	for ip, name := range FindNameTests {
		got, err := findName(net.ParseIP(ip))
		if err != nil {
			t.Error(err)
		}
		if name != got {
			t.Errorf("expected %s, got %s", name, got)
		}
	}
}
