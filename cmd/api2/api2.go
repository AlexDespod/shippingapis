package main

import (
	"fmt"

	"github.com/AlexDespod/shippingapi/internal/servers"
)

func main() {

	serv := new(servers.API2server)

	serv.Init(":9092")

	fmt.Printf("api2 started \n")

	fmt.Print(serv.Run())

}
