package cepgo

import (
	"encoding/json"
)

type (
	CooreiosService struct {
	}
	CorreiosModel struct {
		UF                       string `json:"uf"`
		Localidade               string `json:"localidade"`
		LogradouroDNEC           string `json:"logradouroDNEC"`
		LogradouroTextoAdicional string `json:"logradouroTextoAdicional"`
		LogradouroTexto          string `json:"logradouroTexto"`
		Bairro                   string `json:"bairro"`
		CEP                      string `json:"cep"`

		Street       string `json:"street"`
		Neighborhood string `json:"neighborhood"`
		State        string `json:"state"`
		City         string `json:"city"`
	}
	CooreiosModelResponse struct {
		Erro     bool            `json:"erro"`
		Mensagem string          `json:"mensagem"`
		Total    int             `json:"total"`
		Dados    []CorreiosModel `json:"dados"`
	}
)

func (c *CooreiosService) Execute(cep string, ch chan<- *CEP, errCh chan<- error) {
	data := map[string]string{
		"endereco": cep,
		"tipoCEP":  "ALL",
	}
	body, err := poster("https://buscacepinter.correios.com.br/app/endereco/carrega-cep-endereco.php", data)
	if err != nil {
		errCh <- err
		return
	}
	var model *CooreiosModelResponse
	if err := json.Unmarshal(body, &model); err != nil {
		errCh <- err
		return
	}
	if model.Erro || len(model.Dados) == 0 {
		errCh <- ErrorUnexpectedResponse
		return
	}
	cepResponse := model.Dados[0]

	cepNumber, err := ParseCEPString(cepResponse.CEP)

	if err != nil {
		errCh <- ErrorUnexpectedResponse
		return
	}

	ch <- &CEP{
		Cep:          cepNumber,
		Street:       cepResponse.LogradouroDNEC,
		Neighborhood: cepResponse.Bairro,
		City:         cepResponse.Localidade,
		State:        cepResponse.UF,
	}

	return

}
