package cepgo

import (
	"encoding/json"
	"fmt"
)

type ServiceBrasilAPO struct {
}

type BrasilAPIModel struct {
	CEP          string `json:"cep"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	State        string `json:"state"`
	City         string `json:"city"`
}

func (c *ServiceBrasilAPO) Execute(cep string, ch chan<- CEP) {
	var model *BrasilAPIModel
	body, err := requester(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep))
	if err != nil {
		return
	}
	if err := json.Unmarshal(body, &model); err != nil {
		return
	}

	ch <- CEP{
		Cep:          model.CEP,
		Street:       model.Street,
		Neighborhood: model.Neighborhood,
		City:         model.City,
		State:        model.State,
	}
}
