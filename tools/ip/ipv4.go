package ip

import (
	"fmt"
	"net"
	"sync"
)

var (
	once     sync.Once = sync.Once{}
	ipv4Addr string
)

func setIPv4Addr() {
	once.Do(func() {
		// 获取本机IPV4地址
		ipv4Addr, _ = getMainIPv4()
	})
}

func GetIPv4Addr() string {
	if len(ipv4Addr) == 0 {
		setIPv4Addr()
	}
	return ipv4Addr
}

// GetLocalIPv4s 获取所有本地非回环IPv4地址
func GetLocalIPv4s() ([]string, error) {
	var ips []string

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, ifs := range interfaces {
		if !isPhysicalInterface(ifs) {
			continue
		}

		addrs, err := ifs.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			if ipNet.IP.IsLoopback() || ipNet.IP.To4() == nil {
				continue
			}

			ips = append(ips, ipNet.IP.String())
		}
	}

	if len(ips) == 0 {
		return nil, fmt.Errorf("no suitable IP address found")
	}

	return ips, nil
}

// getMainIPv4 获取主IP地址（通常是第一个非回环地址）
func getMainIPv4() (string, error) {
	ips, err := GetLocalIPv4s()
	if err != nil {
		return "", err
	}
	return ips[0], nil
}
