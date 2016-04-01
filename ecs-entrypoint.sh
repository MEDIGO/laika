#!/bin/sh -e

# workarround for: https://forums.aws.amazon.com/message.jspa?messageID=671321
export STATSD_HOST=$(curl -s 169.254.169.254/latest/meta-data/local-ipv4)

exec "$@"
