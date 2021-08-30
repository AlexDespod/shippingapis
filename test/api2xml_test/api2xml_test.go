package api2xml_test

import (
	"encoding/xml"
	"net"
	"testing"

	procs "github.com/AlexDespod/shippingapi/internal/processes"
	"github.com/AlexDespod/shippingapi/pkg/seal"
	ts "github.com/AlexDespod/shippingapi/types"
	"github.com/stretchr/testify/assert"
)

var mocMess = []byte("<INXMLAPI2><consignee>Harkov</consignee><consignor>Ujgorod</consignor><Cartons><Dem>344.77</Dem><Dem>500</Dem><Dem>54.1</Dem></Cartons></INXMLAPI2>")

func TestGetMessage(t *testing.T) {

	unmarshaled := new(ts.INXMLAPI2)

	xml.Unmarshal(mocMess, unmarshaled)

	testProc := new(procs.API2XMLProcessing)

	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		t.Errorf("some error %v", err.Error())
	}
	go func() {
		conn, err := net.Dial("tcp", "localhost:8081")
		if err != nil {
			t.Errorf("some error %v", err.Error())
		}
		defer conn.Close()

		conn.Write(seal.Seal(mocMess))

	}()

	conn, err := ln.Accept()

	if err != nil {
		t.Errorf("some error %v", err.Error())
	}

	defer conn.Close()

	testProc.SetConn(&conn)

	res, err := testProc.GetMessage()

	if err != nil {
		t.Errorf("some error %v", err.Error())
	}

	assert.EqualValues(t, unmarshaled.Cartons.Dem[0], res.Cartons.Dem[0])
}
