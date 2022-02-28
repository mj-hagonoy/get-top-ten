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
make build
```

### Running the server
```
make run
```
* HTTP server will run in `localhost:8080`

### For Windows
```
make build-win
make run-win
```

* See `Makefile` for more details

