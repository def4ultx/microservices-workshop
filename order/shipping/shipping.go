package shipping

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	return &Client{httpClient}
}

type ShippingInformation struct {
	Address string `json:"address"`
	Status  string `json:"status"`
}

func (c *Client) GetShippingInfo(id string) (*ShippingInformation, error) {
	url := "http://shipping-api:8080/shipping/" + id
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ShippingInformation
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Body)

	return &result, nil
}
