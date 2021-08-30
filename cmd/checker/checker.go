package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/AlexDespod/shippingapi/internal/checkerapi"
	ts "github.com/AlexDespod/shippingapi/types"
)

func main() {

	marshaledList, err := checkerapi.GetDataContainer()

	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}

	respsList := make(ts.ResponseList, 0, 4)

	for _, el := range marshaledList {

		wg.Add(1)

		go func(el ts.MarshaledData) {

			defer wg.Done()

			resp, err := checkerapi.GetResponse(el)

			if err != nil {
				fmt.Println(err)
				return
			}

			respsList = append(respsList, ts.ApiResponse{Cost: resp, Info: el.Info})

		}(el)
	}

	wg.Wait()

	fmt.Print(respsList)
}
