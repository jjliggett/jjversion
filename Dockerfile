FROM golang:1.20.2-alpine3.16@sha256:50e46c122d9b4e9502be0678b52a807be98f54d3fcab417df986bbc4d210dbf7 AS build

ARG BUILD_VERSION
ARG VERSION=${BUILD_VERSION:-0.0.0}

RUN apk add gcc musl-dev

WORKDIR '/app'
COPY go.mod go.sum ./

RUN go mod download

COPY jjversion.go app_version.go ./
COPY jjvercore/*.go jjvercore/

RUN go build -o jjversion -ldflags "-X main.appVersion=${VERSION}"

FROM alpine:3.17.2@sha256:69665d02cb32192e52e07644d76bc6f25abeb5410edc1c7a81a10ba3f0efb90a
WORKDIR '/repo'
COPY --from=build /app/jjversion /usr/local/bin

CMD ["jjversion"]
