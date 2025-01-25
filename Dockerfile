FROM golang:1.23.5-alpine AS modules

COPY go.mod go.sum /modules/

WORKDIR /modules

RUN go mod download

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build  -o /bin/app ./cmd/

FROM alpine
RUN apk add --no-cache bash
COPY .env /
COPY --from=modules /bin/app /home/app
CMD ["/home/app"]