FROM golang:latest

RUN mkdir /app
ADD . /app/
WORKDIR /app

EXPOSE 50051

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY . /usr/src/app/
CMD go run ./client.go