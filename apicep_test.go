package cepgo_test

import (
	"testing"

	"github.com/victorfernandesraton/cepgo"
)

func TestAPICep(t *testing.T) {
	var stub = &cepgo.ServiceApiCEP{}
	testCases := []struct {
		desc      string
		cep       string
		expect    *cepgo.CEP
		waitError error
	}{
		{
			desc: "should be a valid cep",
			cep:  "41342315",
			expect: &cepgo.CEP{
				Cep:   "41342315",
				State: "BA",
				City:  "Salvador",
			},
			waitError: nil,
		},
		{
			desc: "should be a valid cep",
			cep:  "41342-315",
			expect: &cepgo.CEP{
				Cep:   "41342315",
				State: "BA",
				City:  "Salvador",
			},
			waitError: nil,
		},
		{
			desc:      "should be a valid cep",
			cep:       "41342A",
			expect:    nil,
			waitError: cepgo.ErrorUnexpectedResponse,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			chRes := make(chan *cepgo.CEP)
			chErr := make(chan error)
			go stub.Execute(tC.cep, chRes, chErr)

			select {
			case res := <-chRes:
				if res != nil {
					if res.Cep != tC.expect.Cep || res.State != tC.expect.State || res.City != tC.expect.City {
						t.Fatalf("Expected %v, got %v", tC.expect, res)
					}
				}
			case errorInCh := <-chErr:
				if errorInCh != tC.waitError {
					t.Fatalf("Expected %v, got %v", tC.waitError, errorInCh)
				}
			}
		})
	}
}
