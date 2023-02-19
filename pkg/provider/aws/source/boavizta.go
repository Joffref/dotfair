package source

import (
	"encoding/json"
	"fmt"
	"github.com/dotfair-opensource/dotfair/pkg/probe"
	"net/http"
)

const (
	boaviztaApiUrl = "https://api.boavizta.org/v1/"
	boavizta       = "boavizta"
	ProviderName   = "aws"
)

type Response struct {
	GWP probe.Value `json:"gwp" yaml:"gwp"`
	PE  probe.Value `json:"pe" yaml:"pe"`
	ADP probe.Value `json:"adp" yaml:"adp"`
}

type BoaviztaApiClient struct {
	baseUrl    string
	domain     string
	parameters map[string]string
}

func NewBoaviztaApiClient(domain string, params map[string]string) *BoaviztaApiClient {
	return &BoaviztaApiClient{
		baseUrl:    boaviztaApiUrl,
		domain:     domain,
		parameters: params,
	}
}

func (b *BoaviztaApiClient) GetMetrics() (Response, error) {
	url := b.baseUrl + b.domain + "/?"
	for key, value := range b.parameters {
		url += "&" + key + "=" + value
	}
	resp, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("http status code: %d", resp.StatusCode)
	}
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Response{}, err
	}
	return response, nil
}
