package trustip

import (
	"net"
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
	res = IsContained("192.168.100.50", []string{})
	if res == true {
		t.Errorf("empty list must contain no ip")
	}
}

func TestListContains(t *testing.T) {
	t.Parallel()
	list := []string{
		"192.168.100.50",
		"192.168.100.0/24",
	}
	l, err := NewList(list)
	if err != nil {
		t.Fatal(err)
	}
	ip := "192.168.100.50"
	res := l.Contains(ip)
	if res == false {
		t.Errorf("%s: must be contained in list", ip)
	}
	ip = "192.168.100.51"
	res = l.Contains(ip)
	if res == false {
		t.Errorf("%s: must be contained in list", ip)
	}
	ip = "192.168.101.50"
	res = l.Contains(ip)
	if res == true {
		t.Errorf("%s: must be not contained in list", ip)
	}
	ip = "192.168.101.50//100"
	res = l.Contains(ip)
	if res == true {
		t.Errorf("%s: must be not parsed", ip)
	}
}

func BenchmarkIsContained(b *testing.B) {
	list := []string{
		"10.0.99.10",
		"10.0.254.0/16",
		"10.0.90.0/16",
		"10.0.80.0/16",
		"10.0.70.0/16",
		"10.0.60.0/16",
		"10.0.50.0//16",
		"10.0.40.0/16",
		"192.168.100.50",
		"192.168.100.0/24",
	}
	b.Run("ip", func(bb *testing.B) {
		ip := "192.168.100.50"
		bb.ResetTimer()
		for i := 0; i < bb.N; i++ {
			IsContained(ip, list)
		}
	})
	b.Run("cidr", func(bb *testing.B) {
		ip := "192.168.100.51"
		bb.ResetTimer()
		for i := 0; i < bb.N; i++ {
			if !IsContained(ip, list) {
				bb.Fatal("ip must be contained")
			}
		}
	})
}

func BenchmarkListContains(b *testing.B) {
	list := []string{
		"10.0.99.10",
		"10.0.254.0/16",
		"10.0.90.0/16",
		"10.0.80.0/16",
		"10.0.70.0/16",
		"10.0.60.0/16",
		"10.0.50.0//16",
		"10.0.40.0/16",
		"192.168.100.50",
		"192.168.100.0/24",
	}
	l, _ := NewList(list)
	b.Run("ip", func(bb *testing.B) {
		ip := "192.168.100.50"
		bb.ResetTimer()
		for i := 0; i < bb.N; i++ {
			if !l.Contains(ip) {
				bb.Fatal("ip must be contained")
			}
		}
	})
	b.Run("cidr", func(bb *testing.B) {
		ip := "192.168.100.51"
		bb.ResetTimer()
		for i := 0; i < bb.N; i++ {
			if !l.Contains(ip) {
				bb.Fatal("ip must be contained")
			}
		}
	})
}

func BenchmarkIsPrivate(b *testing.B) {
	b.Run("in private", func(bb *testing.B) {
		ip := "192.168.100.50"
		bb.ResetTimer()
		for i := 0; i < bb.N; i++ {
			net.ParseIP(ip).IsPrivate()
		}
	})
	b.Run("no private", func(bb *testing.B) {
		ip := "200.168.100.50"
		bb.ResetTimer()
		for i := 0; i < bb.N; i++ {
			net.ParseIP(ip).IsPrivate()
		}
	})
}
