package inventory

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

type Product struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Amount int    `json:"amount"`
}

type CartResponse struct {
	Products []Product `json:"products"`
}

func (c *Client) GetCartProducts(id string) ([]Product, error) {
	url := "http://inventory-api:8080/cart/" + id
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result CartResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Body)

	return result.Products, nil
}
