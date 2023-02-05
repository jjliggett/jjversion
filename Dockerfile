FROM golang:1.20.0-alpine3.16@sha256:86eb22dea4141189e5656abd0616880387f0515a8d17ee1ef4b2ac79eb6e4ded AS build

ARG BUILD_VERSION
ARG VERSION=${BUILD_VERSION:-0.0.0}

RUN apk add gcc musl-dev

WORKDIR '/app'
COPY go.mod go.sum ./

RUN go mod download

COPY jjversion.go app_version.go ./
COPY jjvercore/*.go jjvercore/

RUN go build -o jjversion -ldflags "-X main.appVersion=${VERSION}"

FROM alpine:3.17.1@sha256:f271e74b17ced29b915d351685fd4644785c6d1559dd1f2d4189a5e851ef753a
WORKDIR '/repo'
COPY --from=build /app/jjversion /usr/local/bin

CMD ["jjversion"]
