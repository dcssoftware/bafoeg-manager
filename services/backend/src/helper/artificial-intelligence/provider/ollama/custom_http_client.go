package rag

import (
	"net/http"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
)

// authRoundTripper is a custom RoundTripper that adds Authorization header
type authRoundTripper struct {
	rt http.RoundTripper
}

func (a *authRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+configuration.OllamaAPI.ApiKey)
	return a.rt.RoundTrip(req)
}

func CreateCustomOllamaHTTPClient() *http.Client {
	return &http.Client{
		Transport: &authRoundTripper{
			rt: http.DefaultTransport,
		},
	}
}
