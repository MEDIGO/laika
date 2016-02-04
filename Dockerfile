FROM gliderlabs/alpine:3.2

RUN apk-install git go mysql-client
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

# Install Node
# This code was taken from https://github.com/nodejs/docker-node/blob/cfc9c3b0dbcbbd3211aec6d525b80ca5e7d1e9ad/4.1/wheezy/Dockerfile
ENV NODE_VERSION 5.5.0
RUN wget https://nodejs.org/dist/v$NODE_VERSION/node-v$NODE_VERSION-linux-x64.tar.gz
RUN tar -xzf "node-v$NODE_VERSION-linux-x64.tar.gz" -C /usr/local
RUN rm "node-v$NODE_VERSION-linux-x64.tar.gz"

RUN go get github.com/stretchr/testify/require
RUN go get github.com/stretchr/testify/suite

WORKDIR /go/src/github.com/MEDIGO/feature-flag
COPY . /go/src/github.com/MEDIGO/feature-flag

RUN go get .

CMD [“feature-flag”]
