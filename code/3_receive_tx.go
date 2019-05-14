package main

import (
    . "github.com/iotaledger/iota.go/api"
    "github.com/iotaledger/iota.go/trinary"
    "github.com/iotaledger/iota.go/converter"
    "fmt"

)

var endpoint = "https://nodes.devnet.thetangle.org"

// The address we want to fetch all transactions for
var address = trinary.Trytes("ZBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")

// We need a query object containing the address we want to look for
var query = FindTransactionsQuery{Addresses: trinary.Hashes{address}}

func main() {
    api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
    must(err)
    
    // Find Transaction Objects uses the connected node to find transactions based on our query
    transactions, err := api.FindTransactionObjects(query)
    must(err)
    
    for _, tx := range transactions {
        // To get our message back we need to convert the signatureMessageFragment to ASCII
        // We should strip all suffix 9's from the signatureMessageFragment, we use a
        // custom function to do this.
        msg, err := converter.TrytesToASCII(removeSuffixNine(tx.SignatureMessageFragment))
        must(err)
        fmt.Println(tx.Hash, " / ", removeSuffixNine(tx.Tag), " / ", msg)
    }
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}

func removeSuffixNine(frag string) string {
    fraglen := len(frag)
    var firstNonNineAt int
    for i := fraglen - 1; i > 0; i-- {
         if frag[i] != '9' {
             firstNonNineAt = i
             break;
        }
    }
    return frag[:firstNonNineAt+1]
}
