# Shipping Service

Manage shipping information

## Rest API

### Shipping

Method      | URI                           | Description                           |
----------- | ----------------------------- | ------------------------------------- |
`GET`       | */shipping/{orderId}*         | Get shipping information              |
`PUT`       | */shipping/{orderId}*         | Update shipping information           |

#### GET /shipping/{orderId}

Response

```json
{
    "address": "111/11 Bangkok 10101",
    "status": "Pending",
}
```

#### PUT /shipping/{orderId}

Request

```json
{
    "address": "111/11 Bangkok 10101",
    "status": "Completed",
}
```

Response

```json
{
    "address": "111/11 Bangkok 10101",
    "status": "Completed",
}
```

## Kafka Consumer

Consume kafka message from Order service and create new shipping record

Message Structure

```json
{
    "address": "111/11 Bangkok 10101",
    "recipient": "John Doe",
    "products": [
        {
            "id": 123456,
            "name": "Test Product 1",
            "amount": 1
        }
    ]
}
```
