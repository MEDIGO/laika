FROM alpine:latest

COPY bin/laika /
COPY public /public/

ENTRYPOINT ["/laika"]
CMD ["run"]
