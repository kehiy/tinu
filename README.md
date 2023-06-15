# tinu
simple URL shortener with Golang and fiber

## Tinu stands for => tiny URL

`
NOTE: this repo still in developing all of stuffs can change rapidly!
`

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


# Options:
* add new tinu (done) (still changing)
* delete tinu   (bug) (still changing)
* update tinu  (done) (still changing)
* use a tinu   (done) (still changing)
