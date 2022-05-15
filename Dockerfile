#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d
RUN go build -o /go/bin/app

#final stage
FROM alpine:latest
LABEL Name=gotodoweb Version=0.0.1

RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app

ENV GIN_MODE=release
ENV MYSQL_HOST=mysql
ENV MYSQL_DATABASE=gotodoweb
ENV MYSQL_USER=root
ENV MYSQL_PASSWORD=password

EXPOSE 3000
ENTRYPOINT /app
