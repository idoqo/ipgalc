package api

import (
	"fmt"
	"strconv"
	"strings"
)

// IPAddr represents an IPv4 address in ipgalc
type IPAddr struct {
	ip   string
	cidr int

	octets       [4]string
	subnetMask   int
	GroupSize    int
	subnet       string
	subnetOctets [4]string
	err          error
}

// NewIPAddr sets up an IPAddr struct for usage
func NewIPAddr(ip string, cidr int) *IPAddr {
	addr := &IPAddr{}
	block, exists := cidrMap[cidr]
	if !exists {
		addr.err = fmt.Errorf("could not find info for provided CIDR: %d", cidr)
	}
	addr.subnet, addr.err = subnetFromCIDR(cidr)

	return &IPAddr{
		ip:         ip,
		cidr:       cidr,
		subnetMask: block.mask,
		GroupSize:  block.size,
	}
}

func (addr *IPAddr) setOctets() error {
	ipSlice := strings.Split(addr.ip, ".")
	octets := [4]string{}
	if len(ipSlice) != 4 {
		return fmt.Errorf("provided IP address is invalid")
	}
	for i, v := range ipSlice {
		n, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		octets[i] = toBin(n)
	}
	addr.octets = octets
	return nil
}

func subnetFromCIDR(cidr int) (string, error) {
	block, exists := cidrMap[cidr]
	if !exists {
		return "", fmt.Errorf("could not find info for provided CIDR: %d", cidr)
	}
	var a, b, c, d int
	if cidr >= 1 && cidr <= 8 {
		a, b, c, d = block.mask, 0, 0, 0
	} else if cidr >= 9 && cidr <= 16 {
		a, b, c, d = 255, block.mask, 0, 0
	} else if cidr >= 17 && cidr <= 24 {
		a, b, c, d = 255, 255, block.mask, 0
	} else {
		a, b, c, d = 255, 255, 255, block.mask
	}
	return fmt.Sprintf("%d.%d.%d.%d", a, b, c, d), nil
}

// ToBinary returns the binary representation of the given IP address
func (addr *IPAddr) ToBinary() (string, error) {
	if addr.octets[0] == "" {
		if err := addr.setOctets(); err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("%s.%s.%s.%s", addr.octets[0], addr.octets[1], addr.octets[2], addr.octets[3]), nil
}

func toBin(n int) string {
	// not using FormatInt so that I can fill in the spaces with 0,
	// that way, each octet is indeed a string of length 8
	//return strconv.FormatInt(int64(n), 2)
	res := ""
	for i := 0; i <= 7; i++ {
		rem := n % 2
		res = strconv.Itoa(rem) + res
		n = n / 2
	}
	return res
}
