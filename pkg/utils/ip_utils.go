package utils

import (
	"net"
)

// IsPrivateIP checks if an IP is private
func IsPrivateIP(ip string) bool {
	privateIPBlocks := []*net.IPNet{
		{IP: net.IPv4(10, 0, 0, 0), Mask: net.CIDRMask(8, 32)},
		{IP: net.IPv4(172, 16, 0, 0), Mask: net.CIDRMask(12, 32)},
		{IP: net.IPv4(192, 168, 0, 0), Mask: net.CIDRMask(16, 32)},
	}
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	for _, block := range privateIPBlocks {
		if block.Contains(parsedIP) {
			return true
		}
	}
	return false
}
