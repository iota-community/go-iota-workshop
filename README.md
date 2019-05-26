# Go IOTA Workshop

Some simple examples to get you started on developing with IOTA using Go.


### Getting started

Go (1.10+) is required to run these examples.

To start playing with these examples run the following commands:

```bash
git clone https://github.com/iota-community/go-iota-workshop.git
cd go-iota-workshop
go mod download
go run code/1_helloworld.go
```

You should receive a message including the statistics of an IOTA node. This means you can explore and run the other examples.
It is highly recommended to change the seeds and addresses used in these examples to make sure you start with a clean slate.


### Examples included

Here are the examples included:

 - `1_helloworld.go`
 - `2_send_tx.go`
 - `3_receive_tx.go`
 - `4_create_address.go`
 - `5_check_balance.go`
 - `6_send_value.go`
 - `7_send_data.go`
 - `8_receive_data.go`
 - `9_zmq.go`


### Work to be done:

 - Add MAM examples
 - Add more complex application examples
 - Refactor by someone with more Go experience :)


#### Contribution

PRs are welcome on `master`
