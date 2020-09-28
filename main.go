package main

import (
	"fmt"
	"log"

	"github.com/idoqo/ipgalc/api"
)

func main() {
	api.Setup()

	ip := "16.194.49.67"
	cidr := 28

	addr := api.NewIPAddr(ip, cidr)
	bin, err := addr.ToBinary()
	if err != nil {
		log.Fatalf("Error ocurred: %s", err.Error())
	}
	fmt.Printf("%s -> %s\n", ip, bin)
}
