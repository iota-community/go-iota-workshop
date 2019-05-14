package main

import (
    . "github.com/iotaledger/iota.go/api"
    "github.com/iotaledger/iota.go/trinary"
    "github.com/iotaledger/iota.go/converter"
    "github.com/iotaledger/iota.go/bundle"
    "fmt"

)

var endpoint = "https://nodes.devnet.thetangle.org"

// Make sure this seed has some funds, you can do this through the faucet on https://faucet.devnet.iota.org
var seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")

// Some random address including a checksum, if we don't have the seed for this the iota sent here are lost
var address = trinary.Trytes("FEEDTHESHEEPS99999999999999999999999999999999999999999999999999999999999999999999LQLNJTGPC")

const mwm = 9
const depth = 3

func main() {
    // compose a new API instance, we provide no PoW function so this uses remote PoW
    api, err := ComposeAPI(HTTPClientSettings{URI: endpoint})
    must(err)

    // convert a ascii message for the transaction to trytes,if possible
    message, err := converter.ASCIIToTrytes("This should have funds!")
    must(err)

    transfers := bundle.Transfers{
        {
            // must be 90 trytes long (include the checksum)
            Address: address,
            Value: 1,
            Message: message,
            Tag: trinary.Trytes("VALTEST"),
        },
    }
    // We need to pass an options object, since we want to use the defaults it stays empty
    prepTransferOpts := PrepareTransfersOptions{}

    trytes, err := api.PrepareTransfers(seed, transfers, prepTransferOpts)
    must(err)
    
    // We don't want to send if the receiving address has been spent from
    spent, err := api.WereAddressesSpentFrom(transfers[0].Address)
    must(err)

    if spent[0] {
        fmt.Println("recipient address is spent from, aborting transfer")
        return
    }
    
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
