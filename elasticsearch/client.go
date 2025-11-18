package elasticsearch

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	BaseURL    string
	HTTPClient Doer
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    strings.TrimSuffix(baseURL, "/"),
		HTTPClient: http.DefaultClient,
	}
}

func NewClientWithHTTP(baseURL string, httpClient Doer) *Client {
	return &Client{
		BaseURL:    strings.TrimSuffix(baseURL, "/"),
		HTTPClient: httpClient,
	}
}

func (c *Client) Search(index string, body any) ([]byte, error) {
	if c == nil || c.HTTPClient == nil {
		return nil, errors.New("elasticsearch client is not initialized")
	}
	if c.BaseURL == "" {
		return nil, errors.New("ELASTIC_SEARCH_URL is empty")
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	url := c.BaseURL + "/" + strings.TrimPrefix(index, "/") + "/_search"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
