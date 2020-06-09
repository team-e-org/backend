FROM golang:1.14.2-alpine AS dev
COPY . /go/src/app
WORKDIR /go/src/app
RUN apk update && apk upgrade && apk add --no-cache git \
    && go get -u \
    && apk del --purge git \
    && go get github.com/cosmtrek/air

ENV SERVER_PORT=3000
CMD air -d

FROM golang:1.14.2-alpine AS build
COPY . /go/src/app
WORKDIR /go/src/app
RUN apk update && apk upgrade && apk add --no-cache git \
    && go get -u \
    && GOOS=linux GOARCH=amd64 go build -o app .

FROM alpine
WORKDIR /go/src/app
COPY --from=build /go/src/app/app .

ENV SERVER_PORT=3000
CMD ["./app"]
