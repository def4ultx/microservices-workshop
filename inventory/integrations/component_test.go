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
)

func TestAddProduct(t *testing.T) {

	type Product struct {
		Name   string `json:"name"`
		Price  int    `json:"price"`
		Amount int    `json:"amount"`
	}

	product := Product{
		Name:   "iPhone X",
		Price:  2990000,
		Amount: 100,
	}

	buffer := bytes.NewBuffer(nil)
	err := json.NewEncoder(buffer).Encode(&product)
	if err != nil {
		t.Errorf("got err: %v", err)
		return
	}

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

	io.Copy(ioutil.Discard, resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code (%v), got (%v)", http.StatusOK, resp.StatusCode)
	}
}
