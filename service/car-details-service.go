package service

import (
	"encoding/json"
	"fmt"
	"goTut/entity"
	"net/http"
)

type CarDetailsService interface {
	GetDetails() entity.CarDetails
}

var (
	carService       CarService   = NewCarService()
	ownerService     OwnerService = NewOwnerService()
	carDataChannel                = make(chan *http.Response)
	ownerDataChannel              = make(chan *http.Response)
)

type services struct{}

func NewCarDetailsService() CarDetailsService {
	return &services{}
}

func (*services) GetDetails() entity.CarDetails {
	//goroutine call enpoint 1
	go carService.FetchData()
	//goroutine call enpoint 2
	go ownerService.FetchData()
	// goroutine get data from https://myfakeapi.com/api/users/1
	car, _ := getCarData()
	owner, _ := getOwnerData()
	return entity.CarDetails{
		ID:        car.ID,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}
}

func getCarData() (entity.Car, error) {
	r1 := <-carDataChannel
	var car entity.Car
	err := json.NewDecoder(r1.Body).Decode(&car)
	if err != nil {
		fmt.Print(err.Error())
		return car, err
	}
	return car, nil
}

func getOwnerData() (entity.Owner, error) {
	r1 := <-ownerDataChannel
	var owner entity.Owner
	err := json.NewDecoder(r1.Body).Decode(&owner)
	if err != nil {
		fmt.Print(err.Error())
		return owner, err
	}
	return owner, nil
}
