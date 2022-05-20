# frontend prepare stage
FROM node:lts-alpine AS frontend_prepare
WORKDIR /frontend
COPY frontend /frontend
RUN yarn install

# frontend build stage
FROM frontend_prepare AS frontend_build
ENV NODE_ENV=production
RUN yarn build

# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
COPY --from=frontend_build /frontend/public /go/src/app/frontend/public
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go get -d
RUN go build -o /go/bin/app

# final stage
FROM scratch
LABEL Name=gotodoweb Version=0.0.1

ENV GIN_MODE=release
ENV MYSQL_HOST=mysql
ENV MYSQL_DATABASE=gotodoweb
ENV MYSQL_USER=root
ENV MYSQL_PASSWORD=password

COPY --from=builder /go/bin/app /app
EXPOSE 3000
CMD ["./app"]
