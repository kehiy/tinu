

# API docs:

## users

### login

endPoint:
`localhost:3000/login`

method: `POST`

request:
```json
{
    "email":"youremail@example.com",
    "password":"123456789"
}
```

response:
```json
{
    "token": "{YOUR_TOKEN}",
}
```

### signup

endPoint:
`localhost:3000/user/delete`

header: `Authorization` : `Bearer {YOUR_TOKEN}`

method: `DELETE`

request:
```json
{
    "email":"youremail@example.com",
    "password":"123456789"
}
```

response:
```json
{
    "message": "user was deleted successfully",
}
```

### delete user

endPoint:
`localhost:3000/signup`

method: `POST`

request:
```json
{
    "email":"youremail@example.com",
    "password":"123456789"
}
```

response:
```json
{
    "message": "OK",
	"data":    "your data",
}
```

## tinus


### addTinu

endPoint:
`localhost:3000/tinu`

header: `Authorization` : `Bearer {YOUR_TOKEN}`

method: `POST`

request:
```json
{
    "url":"https://example.com"
}
```

response:
```json
{
    "data": {
        "id": "f1f118fbe0c41caa",
        "url": "https://example.com",
        "clicked": 0
    },
    "message": "OK"
}
```

### update

endPoint:
`localhost:3000/tinu`

header: `Authorization` : `Bearer {YOUR_TOKEN}`

method: `PATCH`

request:
```json
{
    "id":"f1f118fbe0c41caa",
    "url":"https://amazon.com"
}
```

response:
```json
{
    "data": {
        "id": "f1f118fbe0c41caa",
        "url": "https://amazon.com",
        "clicked": 0
    },
    "message": "updated successfully"
}
```

### delete

endPoint:
`localhost:3000/tinu`

header: `Authorization` : `Bearer {YOUR_TOKEN}`

method: `DELETE`

request:
```json
{
    "id":"8f6ac8fea4eb45e6"
}
```

response:
```json
{
    "id": "8f6ac8fea4eb45e6",
    "message": "deleted successfully"
}
```
### redirect
also this is the short urls:
`localhost:3000/:id`

### check url
you can check where this short url going like this:
`localhost:3000/c/:id`
