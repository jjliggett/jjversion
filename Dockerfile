FROM golang:1.17.1-alpine3.14@sha256:13919fb9091f6667cb375d5fdf016ecd6d3a5d5995603000d422b04583de4ef9 AS build

ARG BUILD_VERSION
ARG VERSION=${BUILD_VERSION:-0.0.0}

WORKDIR '/app'
COPY src/go.mod ./
COPY src/go.sum ./

RUN go mod download

COPY src ./

RUN go build -ldflags "-X main.appVersion=${VERSION}"

FROM alpine:3.14.2@sha256:e1c082e3d3c45cccac829840a25941e679c25d438cc8412c2fa221cf1a824e6a
WORKDIR '/repo'
COPY --from=build /app/jjversion /usr/local/bin

CMD ["jjversion"]
