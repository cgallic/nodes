package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/amovah/eznode"
)

func main() {
	node1 := eznode.NewChainNode(eznode.ChainNodeData{
		Name: "node 1",
		Url:  "https://bscnode2.quickaudits.org/rpc",
		Limit: eznode.ChainNodeLimit{
			Count: 10,
			Per:   5 * time.Second,
		},
		RequestTimeout: 10 * time.Second,
		Priority:       1,
		Middleware:     nil, // optional
	})

	node2 := eznode.NewChainNode(eznode.ChainNodeData{
		Name: "node 2",
		Url:  "https://bscnode3.quickaudits.org/rpc",
		Limit: eznode.ChainNodeLimit{
			Count: 10,
			Per:   5 * time.Second,
		},
		RequestTimeout: 10 * time.Second,
		Priority:       2,
		Middleware:     nil, // optional
	})

	chain := eznode.NewChain(eznode.ChainData{
		Id: "Ethereum",
		Nodes: []*eznode.ChainNode{
			node1,
			node2,
		},
		CheckTickRate: eznode.CheckTick{
			TickRate:         100 * time.Millisecond,
			MaxCheckDuration: 5 * time.Second,
		},
	})

	createdEzNode := eznode.NewEzNode([]*eznode.Chain{chain})

	// sample http request
	req, _ := http.NewRequest("GET", "/latest-block", nil)
	// target ethereum chain
	// eznode will automatically select the node that has the highest priority
	// then will check the node request rate limit
	// if the node is not responding, eznode will try to recover the request
	// and try to send the request to the another node
	response, _ := createdEzNode.SendRequest("Ethereum", req)
	fmt.Println(response)
}
