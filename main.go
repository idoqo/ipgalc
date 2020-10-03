package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/idoqo/ipgalc/api"
)

var ip string
var prefix int

func setup() {
	flag.StringVar(&ip, "ip", "127.0.0.1", "target IP address")
	flag.IntVar(&prefix, "prefix-bits", 24, "prefix bits (number after the slash")
	flag.Parse()
}

func main() {
	setup()

	subnet, err := api.NewSubnet(ip, prefix)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	broadcast := subnet.BroadcastID()
	networkID := subnet.NetworkID()
	hostSize := subnet.HostSize()

	fmt.Println("Broadcast IP: ", broadcast.Ip)
	fmt.Println("Network ID: ", networkID.Ip)
	fmt.Println("Host Size: ", hostSize)
}
