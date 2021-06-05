# Payment Service

Make customer payment via payment gateway. Choose best payment gateway based on customer payment type.

## Rest API

### Payment

Method      | URI                           | Description                           |
----------- | ----------------------------- | ------------------------------------- |
`POST`      | */payment/charge*             | Charge customer payment               |
`GET`       | */payment/charge/{id}*        | View payment information              |

#### POST /payment/charge

Request

```json
{
    "method": "CreditCard",
    "amount": "10000",
    "creditCard": {
        "number": "4111111111111111",
        "expiryMonth": "03",
        "expiryYear": "2030",
        "cvc": "737",
        "holderName": "John Smith"
    }
}
```

Response

```json
{
    "id": 1234,
    "status": "Success"
}
```

#### GET /payment/charge/{id}

Response

```json
{
    "id": 1234,
    "method": "CreditCard",
    "status": "Success",
    "creditCard": {
        "number": "XXXXXXXXXXXX1111",
        "holderName": "John Smith"
    }
}
```
