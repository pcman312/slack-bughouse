# build stage
FROM golang:1.10.3
ADD . /go/src/github.com/JTurpin/slack-bughouse
RUN pwd
RUN ls -l
RUN ls -l src/
RUN cd /go/src/github.com/JTurpin/slack-bughouse && go build -o /slack-bughouse

# Final container
#FROM debian
#MAINTAINER <jim@jimturpin.com>
#
#COPY --from=build-env /src/slack-bughouse /
EXPOSE 9090

CMD ["/slack-bughouse"]
