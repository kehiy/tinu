# tinu
simple URL shortener with Golang and fiber

## Tinu stands for => tiny URL

`
NOTE: this repo still in developing all of stuffs can change rapidly!
`
## useage:
`cd app`
`go run main.go`


# API docs:

## getAll

endPoint:
`localhost:3000/tinu`

method: `GET`

request:
```json
{}
```

response:
```json
[
    {
        "id": "cd983cd8e98ae3f1",
        "url": "https://google.com",
        "clicked": 0
    },
    //.....
]
```
## getOne

endPoint:
`localhost:3000/tinu/:id`

method: `GET`

request:
```json
{}
```

response:
```json
{
    "data": {
        "id": "cd983cd8e98ae3f1",
        "url": "https://google.com",
        "clicked": 0
    }
}
```

## addTinu

endPoint:
`localhost:3000/tinu`

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

## update

endPoint:
`localhost:3000/tinu`

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

## delete

endPoint:
`localhost:3000/tinu`

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
## redirect
also this is the short urls:
`localhost:3000/:id`

# Options:
* add new tinu (done) (still changing)
* delete tinu  (done) (still changing)
* update tinu  (done) (still changing)
* use a tinu   (done) (still changing)
* user signup (developing)
* user login (developing)
* delete account (developing)
* update account (developing)
