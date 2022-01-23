package cili

import (
	"fmt"

	"github.com/cilium/cilium/pkg/client"
)

func main() {
	c, err := client.NewDefaultClient()
	if err != nil {
			...
	}

	endpoints, err := c.EndpointList()
	if err != nil {
			...
	}

	for _, ep := range endpoints {
			fmt.Printf("%8d %14s %16s %32s\n", ep.ID, ep.ContainerName, ep.Addressing.IPV4, ep.Addressing.IPV6)
	}
