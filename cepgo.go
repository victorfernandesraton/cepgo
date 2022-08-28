package cepgo

// Get returns the information about the givem zipcode by concurrent request
func Get(cep string) (*CEP, error) {
	return New().GetCEP(cep)
}
