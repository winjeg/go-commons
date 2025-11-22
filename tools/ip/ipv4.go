package ip

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	once       sync.Once = sync.Once{}
	ipv4Addr   string
	lock       = sync.Mutex{}
	publicIP   string
	lastUpdate = time.Now()
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

// 获取公网IP的函数
func getPublicIP() (string, error) {
	// 定义多个IP查询服务，增加可靠性
	services := []string{
		"https://api.ipify.org",
		"https://ident.me",
		"https://ifconfig.me/ip",
		"https://ipecho.net/plain",
		"https://myexternalip.com/raw",
	}

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	// 尝试每个服务，直到成功获取
	for _, service := range services {
		resp, err := client.Get(service)
		if err != nil {
			continue // 失败则尝试下一个
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				continue
			}
			return string(body), nil
		}
	}
	return "", fmt.Errorf("无法从任何服务获取公网IP")
}

func PublicIPv4() string {
	if len(publicIP) == 0 || lastUpdate.Add(time.Minute).Before(time.Now()) {
		lock.Lock()
		publicIP, _ = getPublicIP()
		lastUpdate = time.Now()
		lock.Unlock()
	}
	return publicIP
}
