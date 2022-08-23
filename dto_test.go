package cepgo_test

import (
	"reflect"
	"testing"

	"github.com/victorfernandesraton/cepgo"
)

func TestValidCep(t *testing.T) {
	tests := []struct {
		cep     string
		expect  string
		wantErr error
	}{
		{
			cep:     "89010220",
			expect:  "89010220",
			wantErr: nil,
		},
		{
			cep:     "890102-20",
			expect:  "89010220",
			wantErr: nil,
		},
		{
			cep:     "89.010-220",
			expect:  "89010220",
			wantErr: nil,
		},
		{
			cep:     "89.010.220",
			expect:  "89010220",
			wantErr: nil,
		},
		{
			cep:     "89.010.22",
			expect:  "",
			wantErr: cepgo.ErrorInvalidCepString,
		},
		{
			cep:     "89.010.229999",
			expect:  "",
			wantErr: cepgo.ErrorInvalidCepString,
		},
		{
			cep:     "89.A10.220",
			expect:  "",
			wantErr: cepgo.ErrorInvalidCepString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.cep, func(t *testing.T) {

			got, err := cepgo.ParseCEPString(tt.cep)
			if err != tt.wantErr {
				t.Errorf("CEPAPIViaCep.GetCEP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("CEPAPIViaCep.GetCEP() = %v, want %v", got, tt.expect)
			}
		})
	}
}
