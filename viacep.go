package cepgo

import (
	"encoding/json"
	"fmt"
)

type (
	ViaCEP struct{}

	ViaCepModel struct {
		CEP         string `json:"cep"`
		Logradouro  string `json:"logradouro"`
		Complemento string `json:"complemento"`
		Bairro      string `json:"bairro"`
		Localidade  string `json:"localidade"`
		UF          string `json:"uf"`
		IBGE        string `json:"ibge"`
		GIA         string `json:"gia"`
		DDD         string `json:"ddd"`
		SIAFI       string `json:"siafi"`
	}
)

func (c *ViaCEP) Execute(cep string, ch chan<- *CEP, errCh chan<- error) {
	var model *ViaCepModel
	body, err := requester(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		errCh <- err
		return
	}
	if err := json.Unmarshal(body, &model); err != nil {
		errCh <- err
		return
	}

	ch <- &CEP{
		Cep:          model.CEP,
		Street:       model.Logradouro,
		Neighborhood: model.Bairro,
		City:         model.Localidade,
		State:        model.UF,
	}
}
