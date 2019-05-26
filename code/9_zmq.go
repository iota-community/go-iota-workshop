package main

import (
    zmq "github.com/pebbe/zmq4"
    "fmt"
    "strings"
)

func main() {
	client, _ := zmq.NewSocket(zmq.SUB)

    // Make sure the connection is closed after stopping the program
    defer client.Close()

    // Connect to a devnet ZMQ address
	client.Connect("tcp://zmq.devnet.iota.org:5556")

    // Subscribe to both tx and sn (confirmed tx) topics
    client.SetSubscribe("tx")
    client.SetSubscribe("sn")

    // Keep looping for messages
    for {
		msg, _ := client.RecvMessage(0)
		for _, str := range msg {
            // We split on space, the should give us the fields as array
            // Fields per topic are covered in the docmentation:
            // https://docs.iota.org/docs/iri/0.1/references/zmq-events
            words := strings.Fields(str)

            if(words[0] == "tx") {
                fmt.Println("New transaction: ", words[1])
            }
            if(words[0] == "sn") {
                fmt.Println("Confirmed transaction: ", words[2], "for milestone", words[1])
            }
		}

	}
}
