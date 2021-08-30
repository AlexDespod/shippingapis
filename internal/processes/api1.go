package src

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	sl "github.com/AlexDespod/shippingapi/pkg/seal"
	ts "github.com/AlexDespod/shippingapi/types"
)

type API1Processing struct {
	conn     *net.Conn
	dataRecv *ts.INAPI1
	response []byte
}

func (a *API1Processing) SetConn(conn *net.Conn) {

	a.conn = conn
}

func (a *API1Processing) SendJson() error {

	err := a.prepareJson()

	if err != nil {

		return err
	}
	conn := *(a.conn)

	_, err = conn.Write(a.response)

	if err != nil {

		return err
	}

	return nil
}

//prepareJson prepare mock data to be sent
func (a *API1Processing) prepareJson() error {

	resp := &ts.OUTAPI1{Total: 1065.05}

	jsondata, err := json.Marshal(resp)

	if err != nil {

		return err
	}

	a.response = sl.Seal(jsondata)

	return nil

}

//DoSomeProcces is a meaning working with received data
func (a *API1Processing) DoSomeProcces() error {
	fmt.Println("Do some work with received data...")
	time.Sleep(time.Millisecond * 50)
	return nil
}

//GetMessage read a tcp connection and unmarshaling it
func (a *API1Processing) GetMessage() (*ts.INAPI1, error) {

	unmarshaled := new(ts.INAPI1)

	messageR, err := sl.Unseal(*a.conn)

	if err != nil {

		return unmarshaled, err
	}

	err = json.Unmarshal(messageR, unmarshaled)

	if err != nil {

		return unmarshaled, err
	}

	a.dataRecv = unmarshaled

	return unmarshaled, nil

}
