package httpclients

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

const (
	brasilapiendpoint = "https://brasilapi.com.br/api/cep/v1/%s"
	viacependpoint    = "https://viacep.com.br/ws/%s/json/"
)

func Getbrasilapi(cep string) string {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(brasilapiendpoint, cep), nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("err")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(bodyBytes)
}

func Getviacep(cep string) string {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf(viacependpoint, cep), nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("err")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(bodyBytes)
}
