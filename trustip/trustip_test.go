package trustip

import (
	"testing"
)

func TestIsContained(t *testing.T) {
	t.Parallel()
	list := []string{
		"192.168.100.50",
		"192.168.100.0/24",
	}
	ip := "192.168.100.50"
	res := IsContained(ip, list)
	if res == false {
		t.Errorf("%s: must be contained in list", ip)
	}
	ip = "192.168.100.51"
	res = IsContained(ip, list)
	if res == false {
		t.Errorf("%s: must be contained in list", ip)
	}
	ip = "192.168.101.50"
	res = IsContained(ip, list)
	if res == true {
		t.Errorf("%s: must be not contained in list", ip)
	}
	ip = "192.168.101.50//100"
	res = IsContained(ip, list)
	if res == true {
		t.Errorf("%s: must be not parsed", ip)
	}
}
