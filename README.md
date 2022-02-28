# Top Ten Words

Accepts an input text and returns the top 10 most used words with their frequencies.

### API

```
POST /top10
Content-Type: "application/json"

Request Body:
{
    "data": "<text for processing>"
}

Response Body:
{
    "top10" : [
        {
            "rank": 1, 
            "value": {
                "words": ["hello"],
                "frequency": 3
            }
        }
    ]
}
```

### Build
```
go build main.go
```

### Run
```
./main
```
or
```
go run main.go
```



