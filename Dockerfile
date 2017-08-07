FROM golang:1.8
MAINTAINER <jim@jimturpin.com>

RUN go get github.com/jturpin/slack-bughouse

WORKDIR /go/src/github.com/jturpin/slack-bughouse

EXPOSE 9090

CMD ["go", "run", "rand-bug.go"]
