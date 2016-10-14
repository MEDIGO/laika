FROM alpine:latest

COPY bin/laika /usr/local/bin/

ENTRYPOINT ["laika"]
CMD ["run"]
