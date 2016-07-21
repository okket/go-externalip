package main

import (
		"github.com/jackpal/gateway"
		natpmp "github.com/jackpal/go-nat-pmp"
		"fmt"
)

func main() {
	gatewayIP, err := gateway.DiscoverGateway()
	if err != nil {
		// fmt.Print("Couldn't determine default gateway\n")
	    return
	}

	client := natpmp.NewClient(gatewayIP)
	response, err := client.GetExternalAddress()
	if err != nil {
		// fmt.Print("No answer from gateway\n")
	    return
	}

	addr := response.ExternalIPAddress
	fmt.Printf("%d.%d.%d.%d\n", addr[0], addr[1], addr[2], addr[3])
}
