package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined")
	}
}

func Start() {

	sanityCheck()
	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
