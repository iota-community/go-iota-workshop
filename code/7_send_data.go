package main

import (
    . "github.com/iotaledger/iota.go/api"
    "github.com/iotaledger/iota.go/trinary"
    "github.com/iotaledger/iota.go/converter"
    "github.com/iotaledger/iota.go/bundle"
    "strings"
    "fmt"

)

var endpoint = "https://nodes.devnet.thetangle.org"

// We need a dummy seed even though we don't sign, because the API requires a seed to send
var seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")
var address = trinary.Trytes("XBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")

// Our data is very long here, it needs to be split over several transactions, 3 in this case
var data = strings.Repeat("This repeated 100 times! ", 100)

const mwm = 9
const depth = 3

func main() {
    // compose a new API instance, we provide no PoW function so this uses remote PoW
    api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
    must(err)

    // convert a ascii message for the transaction to trytes,if possible
    message, err := converter.ASCIIToTrytes(data)
    must(err)

    transfers := bundle.Transfers{
        {
            // must be 90 trytes long (include the checksum)
            Address: address,
            Value: 0,
            Message: message,
            Tag: trinary.Trytes("DATATEST"),
        },
    }
    // We need to pass an options object, since we want to use the defaults it stays empty
    prepTransferOpts := PrepareTransfersOptions{}

    trytes, err := api.PrepareTransfers(seed, transfers, prepTransferOpts)
    must(err)
    
    // Send the transaction to the tangle using given depth and minimum weight magnitude
    bndl, err := api.SendTrytes(trytes, depth, mwm)
    must(err)

    fmt.Println("\nbroadcasted bundle with tail tx hash: ", bundle.TailTransactionHash(bndl))
    // Check what the bundle looks like on thetangle!
    fmt.Printf("https://devnet.thetangle.org/bundle/%s\n\n", bndl[0].Bundle)
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}
