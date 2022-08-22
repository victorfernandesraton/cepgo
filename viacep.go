package cepgo

import (
	"encoding/json"
	"fmt"
)

type ViaCEP struct{}

func (c *ViaCEP) Execute(cep string, ch chan<- CEP) {
	var model *ViaCepModel
	body, err := requester(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return
	}
	if err := json.Unmarshal(body, &model); err != nil {
		return
	}

	ch <- CEP{
		Cep:          model.CEP,
		Street:       model.Logradouro,
		Neighborhood: model.Bairro,
		City:         model.Localidade,
		State:        model.UF,
	}
}
