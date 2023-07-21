package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

// Utilizar * para não criar outro recurso na memória, e assim alterando o recurso interno
func (v *ViaCEP) SetCep(cep string) {
	v.Cep = cep
	fmt.Println(v.Cep)
}

// Atrelando método ao struct ViaCEP
func (v ViaCEP) EnderecoCompleto() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s", v.Logradouro, v.Complemento, v.Bairro, v.Localidade, v.Uf)
}

func soma(a, b int) (int, error) {
	if a > b {
		return a + b, nil
	}

	return 0, errors.New("b não pode ser maior ou igual que a")
}

func main() {
	println("Hello, world")
	//http.ListenAndServe(":8080", nil)

	cep := ViaCEP{
		Cep:         "05187-010",
		Logradouro:  "Av Alexios Jafet 1000",
		Complemento: "X",
		Localidade:  "Jd Ipanema",
		Uf:          "SP",
	}

	fmt.Println(cep.Cep)

	resultado, err := soma(5, 2)

	fmt.Println(resultado, err)

	req, err := http.Get("https://viacep.com.br/ws/" + cep.Cep + "/json/")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close() // fechar conexão (com defer, ele fecha após a sua utilização)

	res, err := io.ReadAll(req.Body) // Le os dados retornados em body
	if err != nil {
		panic(err)
	}

	var data ViaCEP
	err = json.Unmarshal(res, &data) // Deserializa json para string
	if err != nil {
		panic(err)
	}

	fmt.Println(data.Localidade)

	fmt.Println(data.EnderecoCompleto())

	data.SetCep("05100-000")
	fmt.Println(data.Cep)
}
