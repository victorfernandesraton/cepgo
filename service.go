package cepgo

import "errors"

type ServiceRequester interface {
	Execute(cep string, ch chan<- *CEP, errCh chan<- error)
}

type Service struct {
	Providers []ServiceRequester
}

type ServiceResponse struct {
	Cep *CEP
	Err error
}

func (s *Service) ExecuteRequest(cep string) (*CEP, error) {
	var erros []error
	ch := make(chan *CEP)
	err := make(chan error)
	for _, handler := range s.Providers {
		go handler.Execute(cep, ch, err)
	}

	for i := 0; i < len(s.Providers); i++ {
		select {
		case res := <-ch:
			if res != nil {
				return res, nil
			}
		case errorInCh := <-err:
			if errorInCh != nil {
				erros = append(erros, errorInCh)
			}
		}
	}

	if len(erros) == len(s.Providers) {
		return nil, errors.New("none has been conclued")
	}

	return nil, errors.New("none has been conclued")
}
