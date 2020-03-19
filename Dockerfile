FROM golang:1.13

WORKDIR /app

COPY . /app

RUN go mod download

# RUN go run cmd/migration/main.go up

ENTRYPOINT go run cmd/api/main.go