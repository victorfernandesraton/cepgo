package cepgo

type ServiceRequester interface {
	Execute(cep string, ch chan<- CEP)
}

type Service struct {
	Providers []ServiceRequester
}

func (s *Service) ExecuteRequest(cep string) *CEP {
	ch := make(chan CEP)
	for _, handler := range s.Providers {
		go handler.Execute(cep, ch)
	}

	cepResult := <-ch
	return &cepResult
}
