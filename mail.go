package cili

import (
	"fmt"

	"github.com/cilium/cilium/pkg/client"
)

func main() {
	c, err := client.NewDefaultClient()
	if err != nil {
		panic(err)
	}

	endpoints, err := c.EndpointList()
	if err != nil {
		panic(err)
	}

	for _, ep := range endpoints {
		fmt.Printf("%8d %14s %16s %32s\n", ep.ID, ep.Status.ExternalIdentifiers.ContainerName, ep.Status.Networking.HostAddressing.IPV4.AddressType, ep.Status.Networking.HostAddressing.IPV6.AddressType)
	}
}
