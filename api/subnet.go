package api

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Subnet wraps all info regarding a subnet given an IP and n prefix bits
type Subnet struct {
	targetIP   *IPAddr
	subnetMask *IPAddr
	minHost    *IPAddr
	maxHost    *IPAddr
	hostSize   int
	groupSize  int
	prefixBits int
}

func NewSubnet(ip string, prefixBits int) (*Subnet, error) {
	ips, err := splitIP(ip)
	if err != nil {
		return nil, err
	}
	target := &IPAddr{
		Ip:         ip,
		PrefixBits: prefixBits,
		Octets:     ips,
	}

	nmo, err := resolveSubnet(prefixBits)
	if err != nil {
		return nil, err
	}
	subnet := &IPAddr{
		Ip:         fmt.Sprintf("%d.%d.%d.%d", nmo[0], nmo[1], nmo[2], nmo[3]),
		PrefixBits: prefixBits,
		Octets:     nmo,
	}
	return &Subnet{
		targetIP:   target,
		subnetMask: subnet,
	}, nil
}

func splitIP(ip string) (octets [4]int, err error) {
	ipSlice := strings.Split(ip, ".")
	if len(ipSlice) != 4 {
		err = errors.New("invalid IP address")
		return
	}
	for i, str := range ipSlice {
		var n int
		n, err = strconv.Atoi(str)
		if err != nil {
			return
		}
		octets[i] = n
	}
	return
}

// NetworkID produces the network ID of the subnet as a pointer to IPAddr struct
func (sn *Subnet) NetworkID() *IPAddr {
	network := &IPAddr{}

	for i, ipo := range sn.targetIP.Octets {
		network.Octets[i] = ipo & sn.subnetMask.Octets[i]
	}
	network.Ip = fmt.Sprintf("%d.%d.%d.%d", network.Octets[0], network.Octets[1], network.Octets[2], network.Octets[3])
	return network
}

// BroadcastID returns an IPAddr value that holds the broadcast ID of the subnet.
func (sn *Subnet) BroadcastID() *IPAddr {
	broadcast := &IPAddr{}
	for i, ipo := range sn.targetIP.Octets {
		n := sn.subnetMask.Octets[i]
		broadcast.Octets[i] = 256 + (ipo | (^n))
	}
	broadcast.Ip = fmt.Sprintf("%d.%d.%d.%d", broadcast.Octets[0], broadcast.Octets[1], broadcast.Octets[2], broadcast.Octets[3])
	return broadcast
}

// HostSize returns the number of valid hosts possible in the subnet
func (sn *Subnet) HostSize() int {
	n := float64(32 - sn.targetIP.PrefixBits)
	return int(math.Pow(2, n)) - 2
}

func resolveSubnet(prefixBits int) (netmask [4]int, err error) {
	var sn int
	var pb = prefixBits

	if pb == 8 || pb == 16 || pb == 24 || pb == 32 {
		sn = 255
	} else if pb == 7 || pb == 15 || pb == 23 || pb == 31 {
		sn = 254
	} else if pb == 6 || pb == 14 || pb == 22 || pb == 30 {
		sn = 252
	} else if pb == 5 || pb == 13 || pb == 21 || pb == 29 {
		sn = 248
	} else if pb == 4 || pb == 12 || pb == 20 || pb == 28 {
		sn = 240
	} else if pb == 3 || pb == 11 || pb == 19 || pb == 27 {
		sn = 224
	} else if pb == 2 || pb == 10 || pb == 18 || pb == 26 {
		sn = 192
	} else if pb == 1 || pb == 9 || pb == 17 || pb == 25 {
		sn = 128
	}

	if pb >= 1 && pb <= 8 {
		netmask = [4]int{sn, 0, 0, 0}
	} else if pb > 9 && pb <= 16 {
		netmask = [4]int{255, sn, 0, 0}
	} else if pb > 16 && pb <= 24 {
		netmask = [4]int{255, 255, sn, 0}
	} else if pb > 24 && pb <= 32 {
		netmask = [4]int{255, 255, 255, sn}
	} else {
		err = errors.New(fmt.Sprintf("could not find appropriate subnet for %d", pb))
	}
	return

}
