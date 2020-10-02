package api

import (
	"strconv"
)

// IPAddr represents an IPv4 address in ipgalc
type IPAddr struct {
	Ip         string
	PrefixBits int

	Octets [4]int
	Bin    string
}

// ToBinary returns the binary representation of the given IP address
func (addr *IPAddr) ToBinary() string {
	if addr.Bin == "" {
		b := ""
		for _, octet := range addr.Octets {
			b = b + "." + toBin(octet)
		}
		addr.Bin = b
	}
	return addr.Bin
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
