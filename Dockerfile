FROM gliderlabs/alpine:3.2

RUN apk-install git go mysql-client
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

WORKDIR /go/src/github.com/medigo/feat-flag
COPY . /go/src/github.com/medigo/feat-flag

CMD [“feat-flag”]
