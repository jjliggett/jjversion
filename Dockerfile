FROM golang:1.17.4-alpine3.14@sha256:c1261b2f0b62eb03d4d17b9b491d7029ec85e3e384d03ab9e0e24add03388f8f AS build

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
