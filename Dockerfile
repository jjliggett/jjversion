FROM golang:1.18rc1-alpine3.14@sha256:4dd4edb0118078b59c061f8bed9b845f6c6d1b19bada430102a22f6572b71f82 AS build

ARG BUILD_VERSION
ARG VERSION=${BUILD_VERSION:-0.0.0}

WORKDIR '/app'
COPY go.mod go.sum ./

RUN go mod download

COPY jjversion.go app_version.go ./
COPY jjvercore/*.go jjvercore/

RUN go build -o jjversion -ldflags "-X main.appVersion=${VERSION}"

FROM alpine:3.15.0@sha256:21a3deaa0d32a8057914f36584b5288d2e5ecc984380bc0118285c70fa8c9300
WORKDIR '/repo'
COPY --from=build /app/jjversion /usr/local/bin

CMD ["jjversion"]
