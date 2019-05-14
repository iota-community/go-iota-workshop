package main

import (
    . "github.com/iotaledger/iota.go/api"
    "github.com/iotaledger/iota.go/trinary"
    "fmt"
)

// Make sure this address has some funds if you want to test this, you can do this through the faucet on https://faucet.devnet.iota.org
var address = trinary.Trytes("MBYBBFONQZPYZYZHSEZJ9EBEBAFHAZKUFSPBM9YOXJUUAMBUCQQABOWFNPEAGXIGMAVWWFZWDCZJGUTBBZYDSALMPA")
var endpoint = "https://nodes.devnet.thetangle.org"

func main() {
	// compose a new API instance
	api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
	must(err)
    
    // GetNewAddress retrieves the first unspent from address through IRI
    // The 100 argument represents only fully confirmed transactions
	balances, err := api.GetBalances(trinary.Hashes{address}, 100)
	must(err)
    fmt.Println("\nBalance:", balances.Balances[0], " - According to milestone", balances.MilestoneIndex)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
