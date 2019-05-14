package main

import (
    . "github.com/iotaledger/iota.go/api"
    "github.com/iotaledger/iota.go/trinary"
    "github.com/iotaledger/iota.go/converter"
    "fmt"
    "bytes"
    "sort"
)

var endpoint = "https://nodes.devnet.thetangle.org"

// The address we want to fetch all transactions for
// This should contain a set of transactions as provided by the previous example
// It won't work if you ran the previous example multiple times for the same address
var address = trinary.Trytes("XBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")

// We need a query object containing the address we want to look for
var query = FindTransactionsQuery{Addresses: trinary.Hashes{address}}

func main() {
	api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
	must(err)
    
    // Find Transaction Objects uses the connected node to find transactions based on our query
    transactions, err := api.FindTransactionObjects(query)
	must(err)
    
    // We need to sort all transactions by index first so we can concatenate them
    sort.Slice(transactions[:], func(i, j int) bool {
        return transactions[i].CurrentIndex < transactions[j].CurrentIndex
    })
    
    // We define a buffer to concatenate all sorted transactions
    var buffer bytes.Buffer

    for _, tx := range transactions {
        // We add the sorted Transaction Signature Message Fragment to the buffer
        buffer.WriteString(tx.SignatureMessageFragment)
    }
    
    // We need to convert the message to ASCII, but before we do that we need to remove 
    // Additional appended 9's we don't need
    msg, err := converter.TrytesToASCII(removeSuffixNine(buffer.String()))
    must(err)

    // We print out the message
    fmt.Println(msg)
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
