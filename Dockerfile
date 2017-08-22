#FROM golang:1.8
FROM golang:1.8-alpine3.6
MAINTAINER <jim@jimturpin.com>

# uncomment for local dev
#COPY slack-bughouse.go /slack-bughouse.go

# uncomment for github deploys
RUN apk update && apk upgrade && \
    apk add --no-cache bash git
RUN go get github.com/jturpin/slack-bughouse
WORKDIR /go/src/github.com/jturpin/slack-bughouse

EXPOSE 9090

CMD ["go", "run", "slack-bughouse.go"]

# For local dev
#CMD ["go", "run", "/slack-bughouse.go"]
