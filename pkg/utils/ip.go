// Package utils ...
// TODO: разобраться с IPv6 и uint64
package utils

import (
	"encoding/binary"
	"net"
	"strings"
)


// IP2Int fn
func IP2Int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

// Int2IP fn
func Int2IP(n uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, n)
	return ip
}

// GetIPRange accepts x.x.x.x-x.x.x.x, x.x.x.x/x
func GetIPRange(s string) ([]string, error) {
	var start uint32
	var end uint32
	if strings.Contains(s, "-") {
		sp := strings.SplitN(s, "-", 2)
		start = IP2Int(net.ParseIP(sp[0]))
		end = IP2Int(net.ParseIP(sp[1]))
	} else {
		ip, ipnet, err := net.ParseCIDR(s)
		if err != nil {
			return nil, err
		}
		mask := IP2Int(net.IP(ipnet.Mask))
		n := IP2Int(ip)
		start = n & mask
		end = start | ^mask
	}
	var rv []string
	for i := start; i <= end; i++ {
		rv = append(rv, Int2IP(i).String())
	}
	return rv, nil
}
