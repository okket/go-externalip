package main

import natpmp "github.com/jackpal/go-nat-pmp"
import "fmt"

func main() {
	client, err := natpmp.NewClientForDefaultGateway()
	if err != nil {
		// fmt.Print("Couldn't determine default gateway\n")
		return
	}

	response, err := client.GetExternalAddress()
	if err != nil {
		// fmt.Print("No answer from gateway\n")
		return
	}

	addr := response.ExternalIPAddress
	fmt.Printf("%d.%d.%d.%d\n", addr[0], addr[1], addr[2], addr[3])
}
