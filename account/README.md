# Account Service

Manage customer's account

## Rest API

Method      | URI                           | Description                           |
----------- | ----------------------------- | ------------------------------------- |
`POST`      | */user/register*              | Create new account                    |
`POST`      | */user/login*                 | Authenticate user                     |
`GET`       | */user/profile*               | View account information              |
`PUT`       | */user/profile*               | Update account information            |

### POST /user/register

Request

```json
{
    "email": "john@futureskill.com",
    "password": "password1",
    "name": "John"
}
```

Response

```json
{
    "message": "OK"
}
```

### POST /user/login

Request

```json
{
    "email": "john@futureskill.com",
    "password": "password1"
}
```

Response

```json
{
    "token": "dj0yJmk9N2pIazlsZk1iTzIx"
}
```

### GET /user/profile

Header

Key             | Value                         |
--------------- | ----------------------------- |
Authentication  | {user token}                  |

Response

```json
{
    "email": "john@futureskill.com",
    "name": "John"
}
```

### PUT /user/profile

Header

Key             | Value                         |
--------------- | ----------------------------- |
Authentication  | {user token}                  |

Request

```json
{
    "name": "Richard"
}
```

Response

```json
{
    "email": "john@futureskill.com",
    "name": "Richard"
}
```
