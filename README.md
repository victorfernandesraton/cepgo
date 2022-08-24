# gocep: Consulta de informaçõessobre um CEP de forma nativa e concorrente

Fazer consulta de CEP em 3 diferentes provedores e consumir a informação de forma concorrente usando a resposta mais eficiente

## Exemplo

```go
package main

func main() {
    data, err := cepgo.Get("41342315")
    fmt.PrintLn(data)
    // {
    // Cep: "41342315",
    // State: "BA",
    // City: "Salvador",
    // Street: "Estr. do Coqueiro Grande , Fazenda Grande II",
    // "Neighborhood": "Cajazeiras"
    // }
}
```