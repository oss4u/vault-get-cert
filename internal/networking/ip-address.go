package networking

import (
	"fmt"
	"net"
)

func GetIpAddresses() ([]string, error) {
	var result []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return []string{}, fmt.Errorf("failed to get interface addresses: %w", err)
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			result = append(result, ipNet.IP.String())
		}
	}
	return result, nil
}
