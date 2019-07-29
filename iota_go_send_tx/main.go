package main

import (
    . "github.com/iotaledger/iota.go/api"
    "github.com/iotaledger/iota.go/trinary"
    "github.com/iotaledger/iota.go/converter"
    "github.com/iotaledger/iota.go/bundle"
    "fmt"

)

var endpoint = "https://nodes.devnet.thetangle.org"

// We need a dummy seed even though we don't sign, because the API requires a seed to send
var seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")
var address = trinary.Trytes("ZBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")

const mwm = 9
const depth = 3

func main() {
    // compose a new API instance, we provide no PoW function so this uses remote PoW
    api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
    must(err)

    // convert a ascii message for the transaction to trytes,if possible
    message, err := converter.ASCIIToTrytes("Hello World!")
    must(err)

    transfers := bundle.Transfers{
        {
            // must be 90 trytes long (include the checksum)
            Address: address,
            Value: 0,
            Message: message,
            Tag: trinary.Trytes("GOTEST"),
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
    fmt.Printf("https://devnet.thetangle.org/transaction/%s\n\n", bundle.TailTransactionHash(bndl))
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}
