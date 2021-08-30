package main

import (
	"fmt"

	"github.com/AlexDespod/shippingapi/internal/servers"
)

func main() {
	serv := new(servers.API1server)

	serv.Init(":9091")

	fmt.Printf("api1 started \n")

	fmt.Print(serv.Run())

}
