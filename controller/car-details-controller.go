package controller

import (
	"encoding/json"
	"goTut/service"
	"net/http"
)

var carDetailsService service.CarDetailsService

type CarDetailsController interface {
	GetCarDetails(respinse http.ResponseWriter, request *http.Request)
}

func NewCarDetailsConroller(service service.CarDetailsService) CarDetailsController {
	carDetailsService = service
	return &controller{}
}

func (*controller) GetCarDetails(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	result := carDetailsService.GetDetails()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
