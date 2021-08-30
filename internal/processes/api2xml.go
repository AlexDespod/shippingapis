package src

import (
	"encoding/xml"
	"fmt"
	"net"
	"time"

	sl "github.com/AlexDespod/shippingapi/pkg/seal"
	ts "github.com/AlexDespod/shippingapi/types"
)

type API2XMLProcessing struct {
	conn     *net.Conn
	dataRecv *ts.INXMLAPI2
	response []byte
}

func (a *API2XMLProcessing) SetConn(conn *net.Conn) {
	a.conn = conn
}

func (a *API2XMLProcessing) SendXML() error {

	err := a.prepareXML()

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

//prepareXML prepare mock data to be sent
func (a *API2XMLProcessing) prepareXML() error {

	resp := &ts.OUTXMLAPI2{Amount: 1065.05}

	xmldata, err := xml.Marshal(resp)

	if err != nil {

		return err
	}

	a.response = sl.Seal(xmldata)

	return nil

}

//DoSomeProcces is a meaning working with received data
func (a *API2XMLProcessing) DoSomeProcces() error {
	fmt.Println("Do some work with received data...")
	time.Sleep(time.Millisecond * 50)
	return nil
}

//GetMessage read a tcp connection and unmarshaling it
func (a *API2XMLProcessing) GetMessage() (*ts.INXMLAPI2, error) {

	unmarshaled := new(ts.INXMLAPI2)

	messageR, err := sl.Unseal(*a.conn)

	if err != nil {

		return unmarshaled, err
	}

	err = xml.Unmarshal(messageR, unmarshaled)

	if err != nil {

		return unmarshaled, err
	}

	a.dataRecv = unmarshaled

	return unmarshaled, nil

}
