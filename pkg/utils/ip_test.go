package utils

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPToInt(t *testing.T) {
	ip := net.ParseIP("192.168.0.106")
	var n uint32 = 3232235626
	assert.Equal(t, IP2Int(ip), n)
}

func TestGetIPRange(t *testing.T) {
	ips, _ := GetIPRange("10.0.0.1/30")
	assert.ElementsMatch(t, ips, []string{"10.0.0.0", "10.0.0.1", "10.0.0.2", "10.0.0.3"})
}
