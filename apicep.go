package cepgo

import (
	"encoding/json"
	"fmt"
)

type ServiceApiCEP struct {
}

type ApiCepModel struct {
	Code       string `json:"code"`
	Address    string `json:"address"`
	District   string `json:"district"`
	City       string `json:"city"`
	State      string `json:"state"`
	StatusText string `json:"statusText"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	Message    string `json:"message"`
}

func (c *ServiceApiCEP) Execute(cep string, ch chan<- *CEP, errCh chan<- error) {
	var model *ApiCepModel
	body, err := requester(fmt.Sprintf("https://ws.apicep.com/cep/%s.json", cep))
	if err != nil {
		errCh <- ErrorUnexpectedResponse
		return
	}
	if err := json.Unmarshal(body, &model); err != nil {
		errCh <- ErrorUnexpectedResponse
		return
	}

	if !model.Ok && model.Status != 200 {
		errCh <- ErrorUnexpectedResponse
		return

	}

	cepNumber, err := ParseCEPString(model.Code)

	if err != nil {
		errCh <- ErrorUnexpectedResponse
		return
	}
	ch <- &CEP{
		Cep:          cepNumber,
		Street:       model.Address,
		Neighborhood: model.District,
		City:         model.City,
		State:        model.State,
	}

}
