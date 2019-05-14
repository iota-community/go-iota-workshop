package main

import (
    . "github.com/iotaledger/iota.go/api"
    "github.com/davecgh/go-spew/spew"
)

var endpoint = "https://nodes.devnet.thetangle.org"

func main() {
	// compose a new API instance
	api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
	must(err)

	nodeInfo, err := api.GetNodeInfo()
	must(err)

    spew.Dump(nodeInfo)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
