package integration

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/winstonjr/goexpert-desafio-otel/internal/dto"
	"github.com/winstonjr/goexpert-desafio-otel/internal/infra/types"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"io"
	"net/http"
	"strings"
)

type ViacepIntegration struct{}

func NewViacepIntegration() *ViacepIntegration {
	return &ViacepIntegration{}
}

func (o *ViacepIntegration) GetCity(ctx context.Context, cep string, resultch chan<- types.Either[string]) {
	client := getHttpClient()
	cepURL := "https://viacep.com.br/ws/" + cep + "/json/"
	//resp, err := client.Get("https://viacep.com.br/ws/" + cep + "/json/")
	req, err := http.NewRequestWithContext(ctx, "GET", cepURL, nil)
	if err != nil {
		resultch <- types.Either[string]{Left: err}
		return
	}
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	resp, err := client.Do(req)
	if err != nil {
		resultch <- types.Either[string]{Left: err}
		return
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		resultch <- types.Either[string]{Left: err}
		return
	}
	resstr := string(res)
	if strings.Contains(resstr, `"erro"`) {
		resultch <- types.Either[string]{Left: errors.New("can not find zipcode")}
		return
	}
	var data dto.ViacepDTO
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
