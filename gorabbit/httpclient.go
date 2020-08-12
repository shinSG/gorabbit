package gorabbit

import (
	"fmt"
	"io"
	"net/http"
)

// HTTPClient httpclient
type HTTPClient struct {
	http.Client
	user     string
	password string
}

// GetHTTPClient Get RabbitMQ Client
func GetHTTPClient(conf map[string]string) *HTTPClient {
	Client := new(HTTPClient)
	if username, err := conf["username"]; err {
		Client.user = username
	}
	if password, err := conf["password"]; err {
		Client.password = password
	}
	return Client
}

// Send do request
func (hc *HTTPClient) Send(method string, url string, body io.Reader, header map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err, "request error")
	}
	req.SetBasicAuth(hc.user, hc.password)
	for key, val := range header {
		req.Header.Add(key, val)
	}
	resp, err := hc.Do(req)
	return resp, err
}

// Get Method GET
func (hc *HTTPClient) Get(url string) (*http.Response, error) {
	resp, err := hc.Send("GET", url, nil, nil)
	return resp, err
}

// Post Method POST
func (hc *HTTPClient) Post(url string, body io.Reader, header map[string]string) (*http.Response, error) {
	resp, err := hc.Send("POST", url, body, header)
	return resp, err
}

// Post Method POST
func (hc *HTTPClient) DELETE(url string, header map[string]string) (*http.Response, error) {
	resp, err := hc.Send("DELETE", url, nil, header)
	return resp, err
}
