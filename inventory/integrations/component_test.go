// +build integration

package integration

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
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

	io.Copy(ioutil.Discard, resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code (%v), got (%v)", http.StatusOK, resp.StatusCode)
	}
}
