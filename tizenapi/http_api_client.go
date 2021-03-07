package tizenapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	defaultHTTPPort                  = "8001"
	defaultHTTPDialTimeout           = 1 * time.Second
	defaultHTTPRequestTimeout        = 10 * time.Second
	defaultHTTPResponseHeaderTimeout = 10 * time.Second
)

type HTTPAPIOption func(*HTTPAPIClient)

func WithHTTPPort(port string) HTTPAPIOption {
	return func(client *HTTPAPIClient) {
		client.port = port
	}
}

func WithHTTPDialTimeout(timeout time.Duration) HTTPAPIOption {
	return func(client *HTTPAPIClient) {
		client.dialTimeout = timeout
	}
}

func WithHTTPRequestHeaderTimeout(timeout time.Duration) HTTPAPIOption {
	return func(client *HTTPAPIClient) {
		client.requestTimeout = timeout
	}
}

func WithHTTPResponseTimeout(timeout time.Duration) HTTPAPIOption {
	return func(client *HTTPAPIClient) {
		client.responseTimeout = timeout
	}
}

type HTTPAPIClient struct {
	host            string
	port            string
	dialTimeout     time.Duration
	requestTimeout  time.Duration
	responseTimeout time.Duration
	httpClient      *http.Client
}

func NewHTTPAPIClient(host string, options ...HTTPAPIOption) *HTTPAPIClient {
	client := &HTTPAPIClient{
		host:            host,
		port:            defaultHTTPPort,
		dialTimeout:     defaultHTTPDialTimeout,
		requestTimeout:  defaultHTTPRequestTimeout,
		responseTimeout: defaultHTTPResponseHeaderTimeout,
	}

	for _, option := range options {
		option(client)
	}

	client.httpClient = &http.Client{
		Timeout: client.requestTimeout,
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, client.dialTimeout)
			},
			ResponseHeaderTimeout: client.responseTimeout,
		},
	}

	return client
}

func (c *HTTPAPIClient) Host() string {
	return c.host
}

func (c *HTTPAPIClient) Port() string {
	return c.port
}

func (c *HTTPAPIClient) DialTimeout() time.Duration {
	return c.dialTimeout
}

func (c *HTTPAPIClient) RequestTimeout() time.Duration {
	return c.requestTimeout
}

func (c *HTTPAPIClient) ResponseTimeout() time.Duration {
	return c.responseTimeout
}

func (c *HTTPAPIClient) IsAvailable() bool {
	connection, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", c.host, c.port), c.dialTimeout)
	if err != nil {
		return false
	}

	defer func() {
		_ = connection.Close()
	}()

	return true
}

func (c *HTTPAPIClient) GetInfo() (GetInfoResponse, error) {
	response, err := c.httpClient.Get(c.generateHTTPAPIServiceURL())
	if err != nil {
		return GetInfoResponse{}, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return GetInfoResponse{}, fmt.Errorf("invalid response status: %d", response.StatusCode)
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return GetInfoResponse{}, err
	}

	getInfoResponse := GetInfoResponse{}
	err = json.Unmarshal(b, &getInfoResponse)
	if err != nil {
		return GetInfoResponse{}, err
	}

	return getInfoResponse, nil
}

func (c *HTTPAPIClient) GetApp(id string) (GetAppResponse, error) {
	response, err := c.httpClient.Get(c.generateAppURL(id))
	if err != nil {
		return GetAppResponse{}, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return GetAppResponse{}, fmt.Errorf("invalid response status: %d", response.StatusCode)
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return GetAppResponse{}, err
	}

	getAppResponse := GetAppResponse{}
	err = json.Unmarshal(b, &getAppResponse)
	if err != nil {
		return GetAppResponse{}, err
	}

	return getAppResponse, nil
}

func (c *HTTPAPIClient) OpenApp(id string) error {
	response, err := c.httpClient.Post(c.generateAppURL(id), "application/json", nil)
	if err != nil {
		return err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %d", response.StatusCode)
	}

	bodyData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var body interface{}

	err = json.Unmarshal(bodyData, &body)
	if err != nil {
		return err
	}

	switch v := body.(type) {
	case bool:
		if !v {
			return fmt.Errorf("invalid response body: %s", string(bodyData))
		}
	case map[string]interface{}:
		if val, key := v["ok"].(bool); !key || !val {
			return fmt.Errorf("invalid response body: %s", string(bodyData))
		}
	default:
		return fmt.Errorf("invalid response body: %s", string(bodyData))
	}

	return nil
}

func (c *HTTPAPIClient) CloseApp(id string) error {
	request, err := http.NewRequest("DELETE", c.generateAppURL(id), nil)
	if err != nil {
		return err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %d", response.StatusCode)
	}

	return nil
}

func (c *HTTPAPIClient) generateHTTPAPIServiceURL() string {
	return fmt.Sprintf(
		"http://%s:%s/api/v2/",
		c.host,
		defaultHTTPPort,
	)
}

func (c *HTTPAPIClient) generateAppURL(id string) string {
	return fmt.Sprintf(
		"http://%s:%s/api/v2/applications/%s",
		c.host,
		defaultHTTPPort,
		id,
	)
}
