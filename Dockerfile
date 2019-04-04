FROM alpine:latest
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*

RUN update-ca-certificates

RUN apk update && \
    apk add ca-certificates
RUN pwd
RUN ls
COPY bin/laika /
COPY dashboard/public /dashboard/public/

ENTRYPOINT ["/laika"]
CMD ["run"]
