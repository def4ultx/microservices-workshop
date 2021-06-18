// +build integration

package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {

	body := []byte(`{
		"name": "iPhone X",
		"price": 2990000,
		"amount": 100
	}`)
	buffer := bytes.NewBuffer(body)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/product", buffer)
	if err != nil {
		t.Errorf("got err: %v", err)
		return
	}

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Errorf("got err: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code (%v), got (%v)", http.StatusOK, resp.StatusCode)
	}

	product := struct {
		ID     int64  `json:"id"`
		Name   string `json:"name"`
		Price  int    `json:"price"`
		Amount int    `json:"amount"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&product)
	if err != nil {
		t.Errorf("got err while reading body: %v", err)
	}

	assert.Equal(t, product.Name, "iPhone X")
	assert.Equal(t, product.Price, 2990000)
	assert.Equal(t, product.Amount, 100)

	io.Copy(ioutil.Discard, resp.Body)
}
