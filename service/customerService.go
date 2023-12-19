package service

import (
	"banking/domain"
	"banking/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (dc DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return dc.repo.FindAll(status)
}

func (dc DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return dc.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {

	return DefaultCustomerService{repository}
}
