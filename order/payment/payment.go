package payment

import (
	"bytes"
	"encoding/json"
	"errors"
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

type CreditCard struct {
	Number      string `json:"number"`
	ExpiryMonth string `json:"expiryMonth"`
	ExpiryYear  string `json:"expiryYear"`
	CVC         string `json:"cvc"`
	HolderName  string `json:"holderName"`
}

type ChargeRequest struct {
	Method     string     `json:"method"`
	Amount     int        `json:"amount"`
	CreditCard CreditCard `json:"creditCard"`
}

type ChargeResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func (c *Client) ChargeCreditCard(amount int, card CreditCard) (int, error) {
	body := ChargeRequest{
		Method:     "CreditCard",
		CreditCard: card,
		Amount:     amount,
	}

	buffer := bytes.NewBuffer(nil)
	err := json.NewEncoder(buffer).Encode(&body)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPost, "http://payment-api:8080/payment/charge", buffer)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result ChargeResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}
	io.Copy(ioutil.Discard, resp.Body)

	if result.Status != "Success" {
		return 0, errors.New("error from payment api")
	}

	return result.ID, nil
}

type PaymentDetail struct {
	ID     int    `json:"id"`
	Method string `json:"method"`
	Status string `json:"status"`
}

func (c *Client) GetPaymentDetail(id int) (*PaymentDetail, error) {
	url := "http://payment-api:8080/payment/charge/" + strconv.Itoa(id)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var detail PaymentDetail
	err = json.NewDecoder(resp.Body).Decode(&detail)
	if err != nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Body)

	return &detail, nil
}
