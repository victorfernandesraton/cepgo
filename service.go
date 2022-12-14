package cepgo

import (
	"errors"
)

var ErrorNotConclued = errors.New("error not conclued")
var ErrorInAllRequests = errors.New("error in all requests")
var ErrorUnexpectedResponse = errors.New("unexpected response from server") //
var deafultProviders = []ServiceRequester{&ServiceBrasilAPi{}, &ViaCEP{}}

type (
	// ServiceRequester is the interface that wraps the Request method. and recives channels for service
	ServiceRequester interface {
		Execute(cep string, ch chan<- *CEP, errCh chan<- error)
	}
	// Service is a struct that contains the request for various providers and delivery concurrent way to request data
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

// OverrideProvider is a function for creating a new Provider but oveerides services with yours custom providers
func OverrideProvider(providers ...ServiceRequester) Provider {
	return &Service{
		Providers: providers,
	}
}

// GetCEp is a concurrent handler to get data from best api. This behavior is similar to Promise.any in javascript
func (s *Service) GetCEP(cep string) (*CEP, error) {
	parsedCep, errInParse := ParseCEPString(cep)
	if errInParse != nil {
		return nil, errInParse
	}
	var erros []error
	ch := make(chan *CEP)
	err := make(chan error)
	for _, handler := range s.Providers {
		go handler.Execute(parsedCep, ch, err)
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
		return nil, ErrorInAllRequests
	}

	return nil, ErrorNotConclued
}
