package networking

import (
	"fmt"
	"net"
)

// GetIpAddresses returns a list of non-loopback IPv4 addresses.
func GetIpAddresses() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, fmt.Errorf("failed to get interface addresses: %w", err)
	}

	var result []string
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && isNonLoopbackIPv4(ipNet) {
			result = append(result, ipNet.IP.String())
		}
	}
	return result, nil
}

// isNonLoopbackIPv4 checks if the given IPNet is a non-loopback IPv4 address.
func isNonLoopbackIPv4(ipNet *net.IPNet) bool {
	return !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil
}
