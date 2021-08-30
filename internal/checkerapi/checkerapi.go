package checkerapi

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net"

	"github.com/AlexDespod/shippingapi/internal/mock"
	sl "github.com/AlexDespod/shippingapi/pkg/seal"
	ts "github.com/AlexDespod/shippingapi/types"
)

const (
	addr = "Kiev"
	dest = "Lviv"
)

var dementions = [3]float64{float64(100), float64(100), float64(100)}

//GetResponse establish tcp connection with services and exchange messages with them , return data recieved from service
func GetResponse(data ts.MarshaledData) (float64, error) {

	conn, err := net.Dial("tcp", data.Info.Adrr)

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	// fmt.Printf("%s\n", data.Data)

	messageW := sl.Seal(data.Data)

	_, err = conn.Write(messageW)

	if err != nil {
		fmt.Println(err)
	}

	messageR, err := sl.Unseal(conn)

	if err != nil {
		fmt.Println(err)
	}

	return UnmarshalResponse(messageR, data.Info.Name)
}

func UnmarshalResponse(data []byte, APIname string) (float64, error) {

	switch APIname {
	case "API1":
		unmarshaled := ts.OUTAPI1{}
		err := json.Unmarshal(data, &unmarshaled)
		if err != nil {
			return 0, err
		}
		return unmarshaled.Total, nil
	case "API2":
		unmarshaled := ts.OUTAPI2{}
		err := json.Unmarshal(data, &unmarshaled)
		if err != nil {
			return 0, err
		}
		return unmarshaled.Amount, nil
	case "API2XML":
		unmarshaled := ts.OUTXMLAPI2{}
		err := xml.Unmarshal(data, &unmarshaled)
		if err != nil {
			return 0, err
		}
		return unmarshaled.Amount, nil
	default:
		return 0, fmt.Errorf("unknown api name")
	}
}

//GetDataContainer use a mock data for marshalling mock messages in json or xml
func GetDataContainer() (ts.DataContainer, error) {
	var dataCont = make(ts.DataContainer, 0, 4)
	var data []byte
	var err error

	for _, el := range mock.APIs {
		switch el.Name {
		case "API1":
			data, err = MarshalAPI1(addr, dest, dementions)
			if err != nil {
				return nil, err
			}
			dataCont = append(dataCont, ts.MarshaledData{Data: data, Info: el})
		case "API2":
			data, err = MarshalAPI2(addr, dest, dementions)
			if err != nil {
				return nil, err
			}
			dataCont = append(dataCont, ts.MarshaledData{Data: data, Info: el})
		case "API2XML":
			data, err = MarshalAPI2XML(addr, dest, dementions)
			if err != nil {
				return nil, err
			}
			dataCont = append(dataCont, ts.MarshaledData{Data: data, Info: el})
		}
	}

	return dataCont, nil

}

func MarshalAPI1(addr, dest string, dem [3]float64) ([]byte, error) {
	obj := ts.INAPI1{ContactAddress: addr, WarehouseAddress: dest, PackageDimensions: dem}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return data, err
}

func MarshalAPI2(addr, dest string, dem [3]float64) ([]byte, error) {
	obj := ts.INAPI2{Consignee: addr, Consignor: dest, Cartons: dem}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return data, err
}
func MarshalAPI2XML(addr, dest string, dem [3]float64) ([]byte, error) {
	obj := ts.INXMLAPI2{
		Consignee: addr,
		Consignor: dest,
		Cartons: ts.Cartons{
			Dem: []ts.Dem{
				ts.Dem(dem[0]),
				ts.Dem(dem[1]),
				ts.Dem(dem[2]),
			},
		},
	}
	data, err := xml.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return data, err
}
