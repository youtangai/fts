# dts
this is Directory Transfer Service package

# server example
```
package main

import (
    "github.com/youtangai/dts"
)

func main() {
    path := "127.0.0.1" // set ip addr
    port := "5050" // set port number
    dir := "/tmp" // set working dir name. 
    server := dts.NewFileTransferServer(dir, host, port)
    err := server.Run()
    if err != nil {
        panic(err)
    }
}
```

# client example

```
package main

import (
    "github.com/youtangai/dts"
)

func main() {
    path := "127.0.0.1" // set server ip addr
    port := "5050" // set server port number
    dir := "work" // set dir what you want to transfer server

    client, err := dts.NewClient(dir, host, port)
    if err != nil {
        panic(err)
    }
    err = client.TransferDir()
    if err != nil {
        panic(err)
    }
}
```

# dts command 
if you want to read detail example, please read ./cmd/dts/*.go.
## install
`go get github.com/youtangai/dts/cmd/dts`
## usage
### server
`dts srv -host=127.0.0.1 -port=5050 dir`
### client
`dts cli -host=127.0.0.1 -port=5050 dir`
