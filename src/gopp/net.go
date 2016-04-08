package gopp

import (
	"net/http"
)

type HttpClient struct {
	c *http.Client
}

func NewHttpClient() *HttpClient {

	return &HttpClient{}
}
