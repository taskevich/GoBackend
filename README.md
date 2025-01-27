# Запуск

Swagger initialization
```shell
swag init --dir ./cmd,./internal/handler/http/v1,./internal/dto --output ./internal/handler/http/docs
```

Docker build
```shell
docker build . -t gobackend
```

Docker run
```shell
docker run -p 3000:3000 gobackend
```
