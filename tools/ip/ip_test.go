package ip

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLocalIPs(t *testing.T) {
	ip := GetIPv4Addr()
	fmt.Println(ip)
	intIP, _ := IPv4ToUint32(ip)
	assert.Equal(t, ip, Uint32ToIPv4(intIP))
	fmt.Println(Uint32ToIPv4(intIP & 0xffff))
	fmt.Println(PublicIPv4())
}
