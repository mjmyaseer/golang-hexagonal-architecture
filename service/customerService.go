package service

import "banking/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (dc DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {

	return dc.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	
	return DefaultCustomerService{repository}
}