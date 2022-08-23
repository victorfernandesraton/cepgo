package cepgo

import (
	"errors"
)

var ErrorNotConclued = errors.New("error not conclued")
var ErrorInAllRequests = errors.New("error in all requests")
var deafultProviders = []ServiceRequester{&ServiceBrasilAPO{}, &ViaCEP{}}

type (
	ServiceRequester interface {
		Execute(cep string, ch chan<- *CEP, errCh chan<- error)
	}
	Service struct {
		Providers []ServiceRequester
	}
	ServiceResponse struct {
		Cep *CEP
		Err error
	}

	Provider interface {
		GetCEP(cep string) (*CEP, error)
	}
)

func New() Provider {
	return &Service{
		Providers: deafultProviders,
	}
}

func OverrideProvider(providers ...ServiceRequester) Provider {
	return &Service{
		Providers: providers,
	}
}

func (s *Service) AppendProvider(provider ServiceRequester) {
	s.Providers = append(s.Providers, provider)
}

func (s *Service) GetCEP(cep string) (*CEP, error) {
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
		return nil, ErrorNotConclued
	}

	return nil, ErrorInAllRequests
}
