FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /grpc2
RUN mkdir -p /grpc2/protos

WORKDIR /grpc2

COPY ./protos/service.pb.go /grpc2/protos
COPY ./main.go /grpc2

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o grpc2 .

CMD ./grpc2
