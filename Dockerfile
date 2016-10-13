FROM alpine:latest

COPY release/laika /usr/local/bin/

ENTRYPOINT ["laika"]
CMD ["run"]
