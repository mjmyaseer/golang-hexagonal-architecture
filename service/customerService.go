package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
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

func (dc DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := dc.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, err
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {

	return DefaultCustomerService{repository}
}
