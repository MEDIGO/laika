FROM golang:1.6-alpine

RUN apk add --update --no-cache git nodejs
RUN npm install -g bower eslint

RUN go get -u github.com/jteeuwen/go-bindata/...

RUN mkdir -p /go/src/github.com/MEDIGO/laika
WORKDIR /go/src/github.com/MEDIGO/laika

COPY .bowerrc /go/src/github.com/MEDIGO/laika/
COPY bower.json /go/src/github.com/MEDIGO/laika/
RUN bower --allow-root install

COPY . /go/src/github.com/MEDIGO/laika

RUN go get .

CMD [“laika”, "run"]
