package shippingapis

import (
	"bytes"
	"fmt"
)

type DataContainer []MarshaledData

type MarshaledData struct {
	Data []byte //data without seal
	Info APIinfo
}

type ResponseList []ApiResponse

func (r ResponseList) String() string {
	var str bytes.Buffer
	for _, el := range r {
		str.WriteString("cost of this ship will de : \n")
		str.WriteString(fmt.Sprintf("	%f from %s \n", el.Cost, el.Info.Name))

	}
	return fmt.Sprintf("data from apis %s : \n", str.Bytes())
}

type ApiResponse struct {
	Cost float64
	Info APIinfo
}

type INAPI1 struct {
	ContactAddress    string     `json:"contact_address"`
	WarehouseAddress  string     `json:"warehouse_address"`
	PackageDimensions [3]float64 `json:"package_dimensions"`
}

type OUTAPI1 struct {
	Total float64 `json:"total"`
}

type INAPI2 struct {
	Consignee string     `json:"consignee"`
	Consignor string     `json:"consignor"`
	Cartons   [3]float64 `json:"cartons"`
}

type OUTAPI2 struct {
	Amount float64 `json:"amount"`
}

type Dem float64

type Cartons struct {
	Dem []Dem
}
type INXMLAPI2 struct {
	Consignee string `xml:"consignee"`
	Consignor string `xml:"consignor"`
	Cartons   Cartons
}

type OUTXMLAPI2 struct {
	Amount float64 `xml:"amount"`
}

type APIsList []APIinfo

type APIinfo struct {
	Name        string
	Adrr        string
	Communicate string
}
