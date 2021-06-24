FROM golang:1.16-alpine

WORKDIR /cronlog

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    PATH=/cronlog:$PATH

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build

