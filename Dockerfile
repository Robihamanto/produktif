FROM golang:1.13

WORKDIR /app

COPY . /app

RUN go mod download

RUN chmod +x /app/run.sh

ENV COME_RUN_MIGRATION=1

ENTRYPOINT /app/run.sh