# cepgo: Consulta de informações sobre um CEP de forma nativa e concorrente

[![Go](https://github.com/victorfernandesraton/cepgo/actions/workflows/go.yml/badge.svg)](https://github.com/victorfernandesraton/cepgo/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/victorfernandesraton/cepgo.svg)](https://pkg.go.dev/github.com/victorfernandesraton/cepgo)

Fazer consulta de CEP em 3 diferentes provedores e consumir a informação de forma concorrente usando a resposta mais eficiente

## Exemplo

```go
package main

import "github.com/victorfernandesraton/cepgo"

func main() {
    // You shold use 41.342-315 or 41342.315
    data, err := cepgo.Get("41342315")
}
```
