FROM alpine:latest
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*

RUN update-ca-certificates

COPY bin/laika /
COPY public /public/

ENTRYPOINT ["/laika"]
CMD ["run"]
