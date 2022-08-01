# Arguments

ARG BUILD_DATE
ARG GO_VERSION=1.18
ARG ALPINE_VERSION=3.16
ARG AUTHORS="Julien Breux <julien.breux@gmail.com>"

# Stage 0, "builder", based on Go, to build and compile app

# @ref https://hub.docker.com/_/golang
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

WORKDIR /builder
COPY go.mod go.sum Makefile ./
COPY internal internal
COPY pkg pkg
COPY cmd cmd
RUN apk --no-cache add make \
    && make build

# Stage 1, based on Alpine, to have only the compiled app, ready for production

# @ref https://hub.docker.com/_/alpine
FROM alpine:${ALPINE_VERSION}

# @ref https://github.com/opencontainers/image-spec/blob/main/annotations.md
LABEL org.opencontainers.image.created=${BUILD_DATE} \
    org.opencontainers.image.authors=${AUTHORS} \
    org.opencontainers.image.url=https://ghcr.io/speedflow/speedflow \
    org.opencontainers.image.documentation=https://docs.speedflow.app/ \
    org.opencontainers.image.source=https://github.com/speedflow/speedflow

COPY --from=builder /builder/bin/speedflow /bin/speedflow

RUN apk --no-cache --update add ca-certificates

ENTRYPOINT [ "/bin/speedflow" ]
