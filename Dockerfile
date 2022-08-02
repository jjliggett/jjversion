FROM golang:1.18.5-alpine3.16@sha256:dda10a0c69473a595ab11ed3f8305bf4d38e0436b80e1462fb22c9d8a1c1e808 AS build

ARG BUILD_VERSION
ARG VERSION=${BUILD_VERSION:-0.0.0}

WORKDIR '/app'
COPY go.mod go.sum ./

RUN go mod download

COPY jjversion.go app_version.go ./
COPY jjvercore/*.go jjvercore/

RUN go build -o jjversion -ldflags "-X main.appVersion=${VERSION}"

FROM alpine:3.16.0@sha256:686d8c9dfa6f3ccfc8230bc3178d23f84eeaf7e457f36f271ab1acc53015037c
WORKDIR '/repo'
COPY --from=build /app/jjversion /usr/local/bin

CMD ["jjversion"]
