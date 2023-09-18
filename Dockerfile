FROM golang:alpine

ARG DIR_PATH=/usr/src/compiler-rinha

RUN go install github.com/cosmtrek/air@latest

WORKDIR $DIR_PATH

COPY go.mod $DIR_PATH
COPY go.sum $DIR_PATH
RUN go mod download

COPY . $DIR_PATH