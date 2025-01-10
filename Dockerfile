FROM golang:1.23.4-alpine@sha256:6c5c9590f169f77c8046e45c611d3b28fe477789acd8d3762d23d4744de69812 AS build

ARG BUILD_VERSION
ARG VERSION=${BUILD_VERSION:-0.0.0}

RUN apk add gcc musl-dev

WORKDIR '/app'
COPY go.mod go.sum ./

RUN go mod download

COPY jjversion.go app_version.go ./
COPY jjvercore/*.go jjvercore/

RUN go build -o jjversion -ldflags "-X main.appVersion=${VERSION}"

FROM alpine:3.21.2@sha256:56fa17d2a7e7f168a043a2712e63aed1f8543aeafdcee47c58dcffe38ed51099
WORKDIR '/repo'
COPY --from=build /app/jjversion /usr/local/bin

CMD ["jjversion"]
