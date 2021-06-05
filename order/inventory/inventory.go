package inventory

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
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

type CartResponse struct {
	Products []Product `json:"products"`
}

func (c *Client) GetCartProducts(id int) ([]Product, error) {
	url := "http://inventory-api:8080/cart/" + strconv.Itoa(id)
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
