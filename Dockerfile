# syntax = docker/dockerfile:1.3-labs

FROM golang:1-alpine as builder
ARG VERSION=0.0.0
WORKDIR /go/src/cryptocurrency_exporter
COPY . .
RUN apk --no-cache add git openssh build-base
RUN go build -ldflags "-X main.version=${VERSION}" -o app .

FROM alpine as production
LABEL maintainer="rluisr" \
  org.opencontainers.image.url="https://github.com/rluisr/cryptocurrency_exporter" \
  org.opencontainers.image.source="https://github.com/rluisr/cryptocurrency_exporter" \
  org.opencontainers.image.vendor="rluisr" \
  org.opencontainers.image.title="cryptocurrency_exporter" \
  org.opencontainers.image.description="prometheus exporter for getting a crypto currency volume/cap and etc" \
  org.opencontainers.image.licenses="WTFPL"
COPY --from=builder /go/src/cryptocurrency_exporter/app /app
ENTRYPOINT ["/app"]

