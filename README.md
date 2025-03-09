# playground-golang

## Run

```bash
go run main.go
```

```bash
curl localhost:8080/200
curl localhost:8080/400
curl localhost:8080/500
```

## Build

```bash
TAG='playground-golang:latest'
docker build -t $TAG .
docker push $TAG
```
