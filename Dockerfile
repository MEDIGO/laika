FROM alpine:latest

RUN apk update && \
    apk add ca-certificates

COPY bin/laika /
COPY public /public/

ENTRYPOINT ["/laika"]
CMD ["run"]
