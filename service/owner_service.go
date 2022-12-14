package service

import (
	"fmt"
	"net/http"
)

type OwnerService interface {
	FetchData()
}

const (
	ownerServiceUrl = "https://myfakeapi.com/api/users/1"
)

type fetchOwnerDataService struct{}

func NewOwnerService() OwnerService {
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData() {
	client := http.Client{}
	fmt.Printf("Fetching the url %s", ownerServiceUrl)

	resp, _ := client.Get(ownerServiceUrl)

	ownerDataChannel <- resp
}
