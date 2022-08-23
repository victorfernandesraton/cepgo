package cepgo

import (
	"errors"
	"strings"
)

var oveerides = []string{"-", ".", "/"}
var ErrorInvalidCepString = errors.New("Invalid Cep string")

type CEP struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
}

func onlyNumbers(c string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, c)
}

func ParseCEPString(cep string) (string, error) {
	var result string
	result = cep
	for _, o := range oveerides {
		result = strings.ReplaceAll(result, o, "")
	}
	if len(onlyNumbers(result)) != 8 {
		return "", ErrorInvalidCepString
	}
	return result, nil
}
