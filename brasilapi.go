package cepgo

import (
	"encoding/json"
	"fmt"
)

type ServiceBrasilAPi struct {
}

type BrasilAPIModel struct {
	CEP          string `json:"cep"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	State        string `json:"state"`
	City         string `json:"city"`
}

func (c *ServiceBrasilAPi) Execute(cep string, ch chan<- *CEP, errCh chan<- error) {
	var model *BrasilAPIModel
	body, err := requester(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep))
	if err != nil {
		errCh <- ErrorUnexpectedResponse
		return
	}
	if err := json.Unmarshal(body, &model); err != nil {
		errCh <- ErrorUnexpectedResponse
		return
	}

	ch <- &CEP{
		Cep:          model.CEP,
		Street:       model.Street,
		Neighborhood: model.Neighborhood,
		City:         model.City,
		State:        model.State,
	}
}
