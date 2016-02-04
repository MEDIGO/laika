FROM gliderlabs/alpine:3.2

RUN apk-install git go mysql-client
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

RUN apk-install nodejs

RUN go get github.com/stretchr/testify/require
RUN go get github.com/stretchr/testify/suite

WORKDIR /go/src/github.com/MEDIGO/feature-flag
COPY . /go/src/github.com/MEDIGO/feature-flag

RUN npm install
RUN ./node_modules/bower/bin/bower install --allow-root

RUN go get .

CMD [“feature-flag”]
