# Order Service

Manage customer's orders

## Rest API

### Order

Method      | URI                           | Description                           |
----------- | ----------------------------- | ------------------------------------- |
`POST`      | /order                        | Create new order                      |
`GET`       | /order/{id}                   | View order information                |
`GET`       | /orders/{userId}              | List customer's orders                |

#### POST /order

Request

```json
{
    "cartId": 123456,
    "userId": 1234,
    "payment": {
        "method": "CreditCard",
        "creditCard": {
            "number": "4111111111111111",
            "expiryMonth": "03",
            "expiryYear": "2030",
            "cvc": "737",
            "holderName": "John Smith"
        }
    }
}
```

Response

```json
{
    "orderId": "12345",
    "status": "Success"
}
```

#### GET /order/{id}

Response

```json
{
    "orderId": "12345",
    "status": "Completed",
    "totalAmount": 2990000,
    "products": [
        {
            "id": 12345,
            "name": "iPhone X",
            "price": 2990000,
            "amount": 1
        }
    ],
    "payment": {
        "id": 1234,
        "method": "CreditCard",
        "status": "Success"
    },
    "shipping": {
        "address": "111/11",
        "status": "Completed"
    }
}
```

#### GET /orders

Response

```json
{
    "orders": [
        {
            "orderId": "12345",
            "status": "Completed",
            "totalAmount": 2990000,
            "products": [
                {
                    "id": 12345,
                    "name": "iPhone X",
                    "price": 2990000,
                    "amount": 1
                }
            ]
        }
    ]
}
```
