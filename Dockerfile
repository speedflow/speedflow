FROM alpine:3.14

COPY speedflow /bin/speedflow

RUN addgroup -g 1000 -S speedflow && \
    adduser -u 1000 -S speedflow -G speedflow && \
    chown speedflow:speedflow /bin/speedflow

USER speedflow:speedflow

ENTRYPOINT ["/bin/speedflow"]
