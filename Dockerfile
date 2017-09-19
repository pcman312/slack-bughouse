FROM alpine:3.6
MAINTAINER <jim@jimturpin.com>

COPY slack-bughouse /slack-bughouse

EXPOSE 9090

CMD ["/slack-bughouse"]
