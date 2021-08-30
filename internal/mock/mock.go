package mock

import ts "github.com/AlexDespod/shippingapi/types"

var APIs ts.APIsList = ts.APIsList{
	ts.APIinfo{Name: "API1", Adrr: "localhost:9091", Communicate: "json"},
	ts.APIinfo{Name: "API2", Adrr: "localhost:9092", Communicate: "json"},
	ts.APIinfo{Name: "API2XML", Adrr: "localhost:9093", Communicate: "xml"},
}
