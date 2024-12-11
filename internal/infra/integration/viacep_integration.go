package integration

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"
	"io"
	"net/http"
	"strings"
)

type ViacepIntegration struct{}

func NewViacepIntegration() *ViacepIntegration {
	return &ViacepIntegration{}
}

type viacepDTO struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (o *ViacepIntegration) GetCity(cep string, resultch chan<- types.Either[string]) {
	client := getHttpClient()
	req, err := client.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		resultch <- types.Either[string]{Left: err}
		return
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		resultch <- types.Either[string]{Left: err}
		return
	}
	resstr := string(res)
	if strings.Contains(resstr, `"erro"`) {
		resultch <- types.Either[string]{Left: errors.New("can not find zipcode")}
		return
	}
	var data viacepDTO
	err = json.Unmarshal(res, &data)
	if err != nil {
		resultch <- types.Either[string]{Left: err}
		return
	}
	resultch <- types.Either[string]{Right: data.Localidade}
	close(resultch)
}

func getHttpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}
