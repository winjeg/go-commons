package ip

import (
	"fmt"
	"net"
	"strings"
)

func isPhysicalInterface(netInterface net.Interface) bool {
	// 基本过滤
	if netInterface.Flags&net.FlagUp == 0 || netInterface.Flags&net.FlagLoopback != 0 {
		return false
	}

	name := strings.ToLower(netInterface.Name)

	// 明确排除虚拟接口
	virtualKeywords := []string{
		"docker", "veth", "br-", "virbr", "vmnet", "vboxnet",
		"tun", "tap", "wg", "kube", "cni", "flannel",
		"lo", "utun", "awdl", "llw", "virtual", "vethernet",
	}

	for _, keyword := range virtualKeywords {
		if strings.Contains(name, keyword) {
			return false
		}
	}

	// MAC地址过滤
	mac := strings.ToLower(netInterface.HardwareAddr.String())
	virtualMACs := []string{"00:05:69", "00:0c:29", "00:50:56", "02:42:", "0a:00:27"}
	for _, vm := range virtualMACs {
		if strings.HasPrefix(mac, vm) {
			return false
		}
	}
	return true
}

// IPv4ToUint32 IPv4字符串转uint32
func IPv4ToUint32(ipStr string) (uint32, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0, fmt.Errorf("无效的IP地址: %s", ipStr)
	}

	ip = ip.To4()
	if ip == nil {
		return 0, fmt.Errorf("不是IPv4地址: %s", ipStr)
	}

	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3]), nil
}

// Uint32ToIPv4 uint32转IPv4字符串
func Uint32ToIPv4(ipUint uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ipUint>>24),
		byte(ipUint>>16),
		byte(ipUint>>8),
		byte(ipUint),
	)
}
