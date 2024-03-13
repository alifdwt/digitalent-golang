# Digitalent Kominfo Golang 004 - Assignment 2

> This is order management API, built using the [Go programming language](https://golang.org), [Gin](https://gin-gonic.com/), and [Gorm](https://gorm.io/). It allows user to create, read, update, and delete orders.

## 🧰 Installation

1. Clone the repository

```bash
git clone -b assignment-2 https://github.com/alifdwt/digitalent-golang.git
```

2. Install Go and PostgreSQL

3. Run the application

```bash
go run main.go
```

By default, app will be running at localhost:8080

## 📝 API Documentation

You can also access Swagger documentation at http://localhost:8080/docs/index.html

### Create Order

- Method: POST
- Path: /orders
- Request:

```
{
        "orderedAt": "2021-10-06T16:53:27.675931+07:00",
        "customerName": "Test",
        "items":[
            {
                "itemCode": "Item1",
                "description": "ItemDescription",
                "quantity": 1
            },
            {
                "itemCode": "Item2",
                "description": "ItemDescription",
                "quantity": 1
            }
        ]
    }
```

- Response

```
{
	"id": 1,
        "orderedAt": "2021-10-06T16:53:27.675931+07:00",
        "customerName": "Test",
        "items":[
            {
                "itemCode": "Item1",
                "description": "ItemDescription",
                "quantity": 1
            },
            {
                "itemCode": "Item2",
                "description": "ItemDescription",
                "quantity": 1
            }
        ]
    }
```

### Get All Orders

- Method: GET
- Path: /orders
- Response

```
[
	    {
	"id": 1,
        "orderedAt": "2021-10-06T16:53:27.675931+07:00",
        "customerName": "Test",
        "items":[
            {
                "itemCode": "Item1",
                "description": "ItemDescription",
                "quantity": 1
            },
            {
                "itemCode": "Item2",
                "description": "ItemDescription",
                "quantity": 1
            }
        ]
    },
	    {
	"id": 2,
        "orderedAt": "2021-10-06T16:53:27.675931+07:00",
        "customerName": "Test",
        "items":[
            {
                "itemCode": "Item1",
                "description": "ItemDescription",
                "quantity": 1
            },
            {
                "itemCode": "Item2",
                "description": "ItemDescription",
                "quantity": 1
            }
        ]
    }
]
```

### Update order

- Method: PUT
- Path: /orders/{orderId}
- Request:

```
{
    "orderedAt": "2021-10-06T16:53:27.675931+07:00",
    "customerName": "Test update",
    "items":[
        {
            "itemCode": "Item2",
            "description": "ItemDescription",
            "quantity": 1
        },
        {
            "itemCode": "Item2",
            "description": "ItemDescription",
            "quantity": 1
        }
    ]
}
```

- Response

```
{
	"id": 2,
        "orderedAt": "2021-10-06T16:53:27.675931+07:00",
        "customerName": "Test update",
        "items":[
            {
                "itemCode": "Item2",
                "description": "ItemDescription",
                "quantity": 1
            },
            {
                "itemCode": "Item2",
                "description": "ItemDescription",
                "quantity": 1
            }
        ]
}
```

### Delete order

- Method: DELETE
- Path: /orders/{orderId}
- Response

```
"Success delete"
```
