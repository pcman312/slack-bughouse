FROM ubuntu:latest
ADD . /
RUN apt update && apt install -y curl
ENTRYPOINT /test-it.sh
