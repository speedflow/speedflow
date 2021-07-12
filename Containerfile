# Builder

FROM golang:1.16-alpine3.14 AS builder

WORKDIR /builder
COPY go.mod go.sum Makefile ./
COPY internal internal
COPY pkg pkg
COPY cmd cmd
RUN apk --no-cache add make \
    && make build

# Final image

FROM alpine:3.14.0

COPY --from=builder /builder/bin/speedflow /bin/speedflow
RUN apk add --update ca-certificates \
    && rm /var/cache/apk/*
ENTRYPOINT [ "/bin/speedflow" ]
