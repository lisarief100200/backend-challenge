package api

import (
	api "backend-challenge/api/shopping/repositories"
	req "backend-challenge/models"
	"backend-challenge/shop"
	"encoding/json"
	"fmt"
)

type CustomerUsecase interface {
	FindAll() ([]shop.Customer, error)
	Create(customerRequest req.CustomerRequest) (shop.Customer, error)
	Update(ID int, customerRequest req.CustomerRequest) (shop.Customer, error)
	Delete(ID int) (shop.Customer, error)
}

type customerUsecase struct {
	customerRepository api.CustomerRepository
}

func NewCustomerUsecase(customerRepository api.CustomerRepository) *customerUsecase {
	return &customerUsecase{customerRepository}
}

func (u *customerUsecase) FindAll() ([]shop.Customer, error) {
	customers, err := u.customerRepository.FindAll()
	return customers, err
}

func (u *customerUsecase) Create(customerRequest req.CustomerRequest) (shop.Customer, error) {
	total_buy, _ := customerRequest.TotalBuy.Int64()
	creator_id, _ := customerRequest.CreatorId.Int64()

	customers := shop.Customer{
		CustomerName:    customerRequest.CustomerName,
		CustomerContNo:  customerRequest.CustomerContNo,
		CustomerAddress: customerRequest.CustomerAddress,
		TotalBuy:        int(total_buy),
		CreatorId:       int(creator_id),
	}
	newCustomer, err := u.customerRepository.Create(customers)
	return newCustomer, err
}

func (u *customerUsecase) Update(ID int, customerRequest req.CustomerRequest) (shop.Customer, error) {
	customers, err := u.customerRepository.FindByID(ID)
	if err != nil {
		return shop.Customer{}, err
	}

	total_buy, _ := customerRequest.TotalBuy.Int64()
	creator_id, _ := customerRequest.CreatorId.Int64()

	if customerRequest.CustomerName != "" {
		customers.CustomerName = customerRequest.CustomerName
	}
	if total_buy != 0 {
		customerRequest.TotalBuy = json.Number(fmt.Sprintf("%d", total_buy))
	}
	if creator_id != 0 {
		customerRequest.CreatorId = json.Number(fmt.Sprintf("%d", creator_id))
	}
	if customerRequest.CustomerContNo != "" {
		customers.CustomerContNo = customerRequest.CustomerContNo
	}
	if customerRequest.CustomerAddress != "" {
		customers.CustomerAddress = customerRequest.CustomerAddress
	}

	updatedCustomer, err := u.customerRepository.Update(customers)
	if err != nil {
		return shop.Customer{}, err
	}

	return updatedCustomer, nil
}

func (u *customerUsecase) Delete(ID int) (shop.Customer, error) {
	customers, err := u.customerRepository.FindByID(ID)
	if err != nil {
		return shop.Customer{}, err
	}

	newCustomer, err := u.customerRepository.Delete(customers)
	return newCustomer, err
}
