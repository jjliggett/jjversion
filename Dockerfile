FROM golang:alpine AS build

ARG BUILD_VERSION
ARG VERSION=${BUILD_VERSION:-0.0.0}

WORKDIR '/app'
COPY src/go.mod ./
COPY src/go.sum ./

RUN go mod download

COPY src ./

RUN go build -ldflags "-X main.appVersion=${VERSION}"

FROM alpine:latest
WORKDIR '/repo'
COPY --from=build /app/jjversion /usr/local/bin

CMD ["jjversion"]
