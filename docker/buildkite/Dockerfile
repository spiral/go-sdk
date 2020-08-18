FROM golang:1.14

RUN mkdir -p /go/src/github.com/spiral/go-sdk
WORKDIR /go/src/github.com/spiral/go-sdk

ADD go.mod go.sum /go/src/github.com/spiral/go-sdk/
RUN go mod download
