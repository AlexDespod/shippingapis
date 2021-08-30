package main

import (
	"fmt"

	"github.com/AlexDespod/shippingapi/internal/servers"
)

func main() {
	serv := new(servers.API2XMLserver)

	serv.Init(":9093")

	fmt.Printf("api2xml started \n")

	fmt.Print(serv.Run())

}
