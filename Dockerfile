# Download Go dependencies
FROM scrolltech/go-alpine-builder:1.21 as builder

ADD . /chain-monitor
ENV GOPROXY="https://proxy.golang.org,direct"

RUN cd /chain-monitor && \
    go mod tidy &&  \
    go build -v -o ./build/bin/chain-monitor ./cmd/

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest
RUN apk update && apk add vim netcat-openbsd net-tools curl
WORKDIR /app
COPY --from=builder /chain-monitor/build/bin/chain-monitor /bin/
ENTRYPOINT ["chain-monitor"]
