# syntax=docker/dockerfile:1
FROM golang:1.17.5-alpine

ARG BINARY_NAME
ARG HTTP_PORT
ENV BINARY_NAME=$BINARY_NAME
ENV HTTP_PORT=$HTTP_PORT

WORKDIR /$BINARY_NAME

COPY go.mod go.sum .env ./

COPY ./bin ./bin
COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go mod download
RUN go build -o ./bin/$BINARY_NAME ./cmd/server/main.go

EXPOSE $HTTP_PORT

CMD ./bin/start.sh

