package app

import (
	"banking/service"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customers struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode int32  `json:"zip_code" xml:"zip_code"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		err := xml.NewEncoder(w).Encode(customers)
		if err != nil {
			return
		}
	} else {
		err := json.NewEncoder(w).Encode(customers)
		if err != nil {
			return
		}
	}

}
