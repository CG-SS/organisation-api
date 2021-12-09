# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

RUN apk add build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

EXPOSE 12345
RUN go test -v

