package cepgo_test

import (
	"fmt"
	"testing"

	"github.com/victorfernandesraton/cepgo"
)

func TestValidConcurrentSameRequest(t *testing.T) {
	providers := []cepgo.ServiceRequester{&cepgo.ServiceBrasilAPO{}, &cepgo.ViaCEP{}}
	service := &cepgo.Service{
		Providers: providers,
	}
	data := service.ExecuteRequest("41342315")
	fmt.Println(data)
	if data == nil {
		t.Fail()
	}
	if data.State != "BA" {
		t.Fail()
	}

}
