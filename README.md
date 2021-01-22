# URL Shortener
This is a simple URL shortener microservice in GO, you can use Redis or MongoDB

## Settings

### Redis
#### Mac
```sh
export URL_DB=redis
export REDIS_URL=redis://localhost:6379
```
#### Win
```sh
set URL_DB=redis
set REDIS_URL=redis://localhost:6379
```

### MongoDB
#### Mac
```sh
export URL_DB=mongo
export MONGO_URL=mongodb://root:example@localhost:27017/shortener?authSource=admin
export MONGO_DB=shortener
export MONGO_TIMEOUT=30
```
#### Win
```sh
set URL_DB=mongo
set MONGO_URL=mongodb://root:example@localhost:27017/shortener?authSource=admin
set MONGO_DB=shortener
set MONGO_TIMEOUT=30
```

## Run
```sh
cd cmd/cli
go run main.go
```

## Create URL

### POST /
```sh
curl --location --request POST 'http://localhost:8000' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "https://github.com/guiaramos/go-url-shortener"
}'
```

### User the URL
```sh
curl --location --request GET 'http://localhost:8000/5emxNyBMR'
```