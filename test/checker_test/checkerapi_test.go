package checker_test

import (
	"testing"

	"github.com/AlexDespod/shippingapi/internal/checkerapi"
	"github.com/AlexDespod/shippingapi/internal/mock"
	ts "github.com/AlexDespod/shippingapi/types"
	"github.com/stretchr/testify/assert"
)

type par struct {
	addr string
	dest string
	dems [3]float64
}

var mockStructure = []struct {
	Expected   string
	Parameters par
}{
	{
		Expected:   "<INXMLAPI2><consignee>Kiev</consignee><consignor>Lviv</consignor><Cartons><Dem>100</Dem><Dem>100</Dem><Dem>100</Dem></Cartons></INXMLAPI2>",
		Parameters: par{addr: "Kiev", dest: "Lviv", dems: [3]float64{float64(100), float64(100), float64(100)}},
	},
	{
		Expected:   "<INXMLAPI2><consignee>Odessa</consignee><consignor>Jitomir</consignor><Cartons><Dem>10.51</Dem><Dem>100</Dem><Dem>77.65</Dem></Cartons></INXMLAPI2>",
		Parameters: par{addr: "Odessa", dest: "Jitomir", dems: [3]float64{10.51, float64(100), 77.65}},
	},
	{
		Expected:   "<INXMLAPI2><consignee>Harkov</consignee><consignor>Ujgorod</consignor><Cartons><Dem>344.77</Dem><Dem>500</Dem><Dem>54.1</Dem></Cartons></INXMLAPI2>",
		Parameters: par{addr: "Harkov", dest: "Ujgorod", dems: [3]float64{344.77, float64(500), 54.1}},
	},
}

//for this test require one running service at least
func TestGetResponse(t *testing.T) {

	exp := 1065.050000
	info := mock.APIs[2]

	for _, v := range mockStructure {

		mockMarsahledData := ts.MarshaledData{Data: []byte(v.Expected), Info: info}

		res, err := checkerapi.GetResponse(mockMarsahledData)

		if err != nil {
			t.Errorf("some error %v", err.Error())
		}
		assert.EqualValues(t, exp, res)
	}

}

func TestMarshalAPI2XML(t *testing.T) {

	for _, v := range mockStructure {
		res, err := checkerapi.MarshalAPI2XML(v.Parameters.addr, v.Parameters.dest, v.Parameters.dems)
		if err != nil {
			t.Errorf("some error %v", err.Error())
		}
		assert.EqualValues(t, v.Expected, res)
	}

}
