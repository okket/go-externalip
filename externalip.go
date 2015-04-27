package main

import natpmp "github.com/jackpal/go-nat-pmp"
import gateway "github.com/jackpal/gateway"
import "fmt"

func main() {
	gwAddr, err := gateway.DiscoverGateway()
	if err != nil {
		// fmt.Print("Couldn't dertermine default gateway\n")
		return
	}

	client := natpmp.NewClient(gwAddr)
	response, err := client.GetExternalAddress()
	if err != nil {
		// fmt.Printf("No answer from gateway %v\n", gatewayAddr)
		return
	}
	addr := response.ExternalIPAddress
	fmt.Printf("%d.%d.%d.%d\n", addr[0], addr[1], addr[2], addr[3])
}
