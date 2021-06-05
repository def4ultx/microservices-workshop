package payment

import (
	"bytes"
	"encoding/json"
	"errors"
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

type CreditCard struct {
	Number      string `json:"number"`
	ExpiryMonth string `json:"expiryMonth"`
	ExpiryYear  string `json:"expiryYear"`
	CVC         string `json:"cvc"`
	HolderName  string `json:"holderName"`
}

type ChargeRequest struct {
	Method     string     `json:"method"`
	CreditCard CreditCard `json:"creditCard"`
}

type ChargeResponse struct {
	Status string `json:"status"`
}

func (c *Client) ChargeCreditCard(card CreditCard) error {

	body := ChargeRequest{
		Method:     "CreditCard",
		CreditCard: card,
	}

	buffer := bytes.NewBuffer(nil)
	err := json.NewEncoder(buffer).Encode(&body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://payment:8080/payment/charge", buffer)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result ChargeResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}
	io.Copy(ioutil.Discard, resp.Body)

	if result.Status != "Success" {
		return errors.New("error from payment api")
	}

	return nil
}