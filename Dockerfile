FROM golang:1.6-alpine

RUN apk add --update --no-cache git nodejs curl
RUN npm install -g bower eslint

ENV GLIDE_VERSION 0.11.1
RUN curl -sSL https://github.com/Masterminds/glide/releases/download/v$GLIDE_VERSION/glide-v$GLIDE_VERSION-linux-386.tar.gz -o glide.tar.gz \
    && tar -xzf glide.tar.gz \
    && mv linux-386/glide /usr/local/bin \
    && rm -rf glide.tar.gz linux-386

RUN go get -u github.com/jteeuwen/go-bindata/...

RUN mkdir -p /go/src/github.com/MEDIGO/laika
WORKDIR /go/src/github.com/MEDIGO/laika

COPY glide.yaml /go/src/github.com/MEDIGO/laika/
COPY glide.lock /go/src/github.com/MEDIGO/laika/
RUN glide install

COPY .bowerrc /go/src/github.com/MEDIGO/laika/
COPY bower.json /go/src/github.com/MEDIGO/laika/
RUN bower --allow-root install

COPY . /go/src/github.com/MEDIGO/laika
RUN go install .

CMD [“laika”, "run"]
