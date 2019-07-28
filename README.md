# Go IOTA Workshop

Some simple examples to get you started on developing with IOTA using Go.


### Getting started

Go (1.10+) is required to run these examples.

To start playing with these examples run the following commands:

```bash
git clone https://github.com/iota-community/go-iota-workshop.git
cd go-iota-workshop
go mod download
go run iota_go_helloworld/main.go
```

You should receive a message including the statistics of an IOTA node. This means you can explore and run the other examples.
It is highly recommended to change the seeds and addresses used in these examples to make sure you start with a clean slate.


### Examples included

Here are the examples included:

 - 1: `iota_go_helloworld`
 - 2: `iota_go_send_tx`
 - 3: `iota_go_receive_tx`
 - 4: `iota_go_create_address`
 - 5: `iota_go_check_balance`
 - 6: `iota_go_send_value`
 - 7: `iota_go_send_data`
 - 8: `iota_go_receive_data`
 - 9: `iota_go_zmq`


### Work to be done:

 - Add MAM examples
 - Add more complex application examples
 - Refactor by someone with more Go experience :)


#### Contribution

PRs are welcome on `master`
