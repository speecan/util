package trustip

import (
	"fmt"
	"net"
)

type (
	List struct {
		IPs   []net.IP
		CIDRs []*net.IPNet
	}
)

// IsContained check if IP is within list
func IsContained(ip string, list []string) bool {
	l, _ := NewList(list)
	return l.Contains(ip)
}

func NewList(list []string) (*List, error) {
	l := &List{
		IPs:   make([]net.IP, 0),
		CIDRs: make([]*net.IPNet, 0),
	}
	var reterr error
	for _, v := range list {
		if ip := net.ParseIP(v); ip != nil { // ip
			l.IPs = append(l.IPs, ip)
			continue
		}
		_, cidr, err := net.ParseCIDR(v) // cidr
		if err != nil {
			reterr = fmt.Errorf("%s was neither IP nor CIDR: %w", v, err)
			continue
		}
		l.CIDRs = append(l.CIDRs, cidr)
	}
	return l, reterr
}

func (x *List) Contains(ip string) bool {
	t := net.ParseIP(ip)
	if t == nil {
		return false
	}
	for _, v := range x.CIDRs {
		if v.Contains(t) {
			return true
		}
	}
	for _, v := range x.IPs {
		if v.Equal(t) {
			return true
		}
	}
	return false
}
