package cepgo_test

import (
	"errors"
	"testing"

	"github.com/victorfernandesraton/cepgo"
)

var ErrorSumulated = errors.New("Simulated error")

func TestValidConcurrentSameRequest(t *testing.T) {
	service := &cepgo.Service{
		Providers: []cepgo.ServiceRequester{&cepgo.ServiceBrasilAPO{}, &cepgo.ViaCEP{}},
	}
	data, err := service.ExecuteRequest("41342315")
	if err != nil {
		t.Fail()
	}
	if data == nil {
		t.Fail()
	}
	if data.State != "BA" {
		t.Fail()
	}

}

type CustonError struct {
}

func (c *CustonError) Execute(cep string, ch chan<- *cepgo.CEP, errCh chan<- error) {
	errCh <- ErrorSumulated
	return
}

func TestWithAllErrorInRequest(t *testing.T) {
	service := &cepgo.Service{
		Providers: []cepgo.ServiceRequester{&CustonError{}, &CustonError{}},
	}
	data, err := service.ExecuteRequest("41342315")
	if err == nil {
		t.Fatalf("expect %v, got %v", ErrorSumulated, err)
	}
	if data != nil {
		t.Fatalf("expect %v, got %v", true, data)
	}
}

func TestWithOneErrorANdOneSucess(t *testing.T) {
	service := &cepgo.Service{
		Providers: []cepgo.ServiceRequester{&cepgo.ServiceBrasilAPO{}, &CustonError{}},
	}
	data, err := service.ExecuteRequest("41342315")
	if err != nil {
		t.Fatalf("expect %v, got %v", ErrorSumulated, err)
	}
	if data == nil {
		t.Fatalf("expect %v, got %v", true, data)
	}
}
