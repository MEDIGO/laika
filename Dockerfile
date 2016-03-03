FROM golang:1.6-alpine

RUN apk add --update --no-cache git mysql-client nodejs

RUN npm install -g gulp bower

ENV GLIDE_VERSION 0.8.3
ENV GLIDE_URL https://github.com/Masterminds/glide/releases/download/$GLIDE_VERSION/glide-$GLIDE_VERSION-linux-amd64.tar.gz
RUN curl -fsSL "$GLIDE_URL" -o glide.tar.gz \
 	&& tar -xzf glide.tar.gz \
 	&& mv linux-amd64/glide /usr/local/bin \
 	&& rm -rf linux-amd64 \
 	&& rm glide.tar.gz

RUN mkdir -p /go/src/github.com/MEDIGO/laika
WORKDIR /go/src/github.com/MEDIGO/laika

COPY glide.lock /go/src/github.com/MEDIGO/laika/
COPY glide.yaml /go/src/github.com/MEDIGO/laika/
RUN glide install

COPY package.json /go/src/github.com/MEDIGO/laika/
RUN npm install

COPY bower.json /go/src/github.com/MEDIGO/laika/
RUN bower --allow-root install

COPY . /go/src/github.com/MEDIGO/laika

RUN go get .

CMD [“laika”]
