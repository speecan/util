package trustip

import (
	"log"
	"net"
)

// IsContained check if IP is within list
func IsContained(ip string, list []string) bool {
	for _, v := range list {
		_, network, err := net.ParseCIDR(v)
		if err == nil {
			ip := net.ParseIP(ip)
			if network.Contains(ip) {
				return true
			}
		} else {
			if net.ParseIP(ip).Equal(net.ParseIP(v)) {
				return true
			}
		}
	}
	log.Println(ip + ": is not trusted")
	return false
}
