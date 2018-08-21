FROM ubuntu:latest
ADD . /
RUN apt update && apt install -y curl
RUN /test-it.sh
