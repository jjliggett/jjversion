FROM golang:1.22.2-alpine@sha256:cdc86d9f363e8786845bea2040312b4efa321b828acdeb26f393faa864d887b0 AS build

ARG BUILD_VERSION
ARG VERSION=${BUILD_VERSION:-0.0.0}

RUN apk add gcc musl-dev

WORKDIR '/app'
COPY go.mod go.sum ./

RUN go mod download

COPY jjversion.go app_version.go ./
COPY jjvercore/*.go jjvercore/

RUN go build -o jjversion -ldflags "-X main.appVersion=${VERSION}"

FROM alpine:3.20.0@sha256:77726ef6b57ddf65bb551896826ec38bc3e53f75cdde31354fbffb4f25238ebd
WORKDIR '/repo'
COPY --from=build /app/jjversion /usr/local/bin

CMD ["jjversion"]
