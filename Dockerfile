FROM gliderlabs/alpine:3.2

RUN apk-install git go mysql-client
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
ENV GLIDE_VERSION 0.8.3
ENV GLIDE_URL https://github.com/Masterminds/glide/releases/download/$GLIDE_VERSION/glide-$GLIDE_VERSION-linux-amd64.tar.gz

RUN apk-install nodejs

RUN curl -fsSL "$GLIDE_URL" -o glide.tar.gz \
 	&& tar -xzf glide.tar.gz \
 	&& mv linux-amd64/glide /usr/local/bin \
 	&& rm -rf linux-amd64 \
 	&& rm glide.tar.gz

RUN go get github.com/stretchr/testify/require
RUN go get github.com/stretchr/testify/suite

WORKDIR /go/src/github.com/MEDIGO/laika
COPY . /go/src/github.com/MEDIGO/laika

COPY glide.lock /go/src/github.com/medigo/core/
COPY glide.yaml /go/src/github.com/medigo/core/
RUN glide install

RUN npm install
RUN ./node_modules/bower/bin/bower install --allow-root

RUN go get .

CMD [“laika”]
