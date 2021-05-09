# Inventory Service

Manages application's products and shopping cart functionality

## Rest API

### Product catalog

Method      | URI                           | Description                           |
----------- | ----------------------------- | ------------------------------------- |
`GET`       | */products/recommendations*   | Get recommended products              |
`GET`       | */products*                   | List products                         |
`POST`      | */product*                    | Add new product                       |
`GET`       | */product/{id}*               | Fetch product information based on id |
`PUT`       | */product/{id}*               | Updates existing product              |

### GET /products/recommendations

Response

```json
{
    "products": [
        {
            "id": 12345,
            "name": "iPhone X",
            "price": 2990000,
            "amount": 100
        }
    ]
}
```

### GET /products/

Response

```json
{
    "products": [
        {
            "id": 12345,
            "name": "iPhone X",
            "price": 2990000,
            "amount": 100
        }
    ]
}
```

### POST /product/

Request

```json
{
    "name": "iPhone X",
    "price": 2990000,
    "amount": 100
}
```

Response

```json
{
    "id": 12345,
    "name": "iPhone X",
    "price": 2990000,
    "amount": 100
}
```

### GET /product/{id}

Response

```json
{
    "id": 12345,
    "name": "iPhone X",
    "price": 2990000,
    "amount": 100
}
```

### PUT /product/{id}

Request

```json
{
    "id": 12345,
    "name": "iPhone X",
    "price": 3990000,
    "amount": 250
}
```

Response

```json
{
    "id": 12345,
    "name": "iPhone X",
    "price": 3990000,
    "amount": 250
}
```

### Shopping Cart

Method      | URI                                   | Description                           |
----------- | ------------------------------------- | ------------------------------------- |
`POST`      | */cart*                               | Create new cart                       |
`GET`       | */cart/{cartId}*                      | Get cart item by id                   |
`DELETE`    | */cart/{cartId}products*              | Remove all item from cart             |
`POST`      | */cart/{cartId}product/{productId}*   | Add item into cart                    |
`DELETE`    | */cart/{cartId}product/{productId}*   | Remove item from cart                 |

#### POST /cart

Request

```json
{
    "userId": 123
}
```

Response

```json
{
    "cartId": 656991997564649473
}
```

#### GET /cart/{cartId}

Response

```json
{
    "products": [
        {
            "id": 656948471224139777,
            "name": "testtest",
            "price": 9900,
            "amount": 99
        },
        {
            "id": 656964013040828417,
            "name": "test product 4",
            "price": 9900,
            "amount": 99
        }
    ]
}
```

#### DELETE /cart/{cartId}/products

Response

```json
{}
```

#### POST /cart/{cartId}/product/{productId}

Request

```json
{
    "amount": 1
}
```

Response

```json
{}
```

#### DELETE /cart/{cartId}/product/{productId}

Response

```json
{}
```
