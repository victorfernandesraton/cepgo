package cepgo_test

import (
	"errors"
	"testing"

	"github.com/victorfernandesraton/cepgo"
)

var ErrorSumulated = errors.New("Simulated error")

func TestValidConcurrentSameRequest(t *testing.T) {

	data, err := cepgo.New().GetCEP("41342315")
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
	data, err := cepgo.OverrideProvider(&CustonError{}, &CustonError{}).GetCEP("41342315")
	if err != cepgo.ErrorInAllRequests {
		t.Fatalf("expect %v, got %v", cepgo.ErrorInAllRequests, err)
	}
	if data != nil {
		t.Fatalf("expect %v, got %v", true, data)
	}
}

func TestWithOneErrorANdOneSucess(t *testing.T) {
	data, err := cepgo.OverrideProvider(&cepgo.ServiceBrasilAPO{}, &CustonError{}).GetCEP("41342315")
	if err != nil {
		t.Fatalf("expect %v, got %v", ErrorSumulated, err)
	}
	if data == nil {
		t.Fail()
	}
	if data.State != "BA" {
		t.Fail()
	}
}
