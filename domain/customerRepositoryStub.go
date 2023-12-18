package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (cr CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return cr.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	return CustomerRepositoryStub{
		customers: []Customer{
			{
				Id:          "1001",
				Name:        "Ashish",
				City:        "Delhi",
				Zipcode:     "0100",
				DateOfBirth: "23-12-2000",
				Status:      "1",
			}, {
				Id:          "1002",
				Name:        "Shukla",
				City:        "Delhi",
				Zipcode:     "0100",
				DateOfBirth: "23-12-2002",
				Status:      "1",
			},
		},
	}
}
