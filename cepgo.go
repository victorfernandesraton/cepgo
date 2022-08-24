package cepgo

func Get(cep string) (*CEP, error) {
	return New().GetCEP(cep)
}
