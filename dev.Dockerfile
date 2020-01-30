FROM golang:1.13-alpine

# Required for running go tests
RUN apk --no-cache add gcc g++ make ca-certificates

WORKDIR /go/src/github.com/MEDIGO/laika

COPY go.mod /go/src/github.com/MEDIGO/laika
COPY go.sum /go/src/github.com/MEDIGO/laika

RUN go mod download

RUN go get github.com/ivpusic/rerun

COPY . /go/src/github.com/MEDIGO/laika

EXPOSE 8000