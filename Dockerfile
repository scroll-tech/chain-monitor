# Download Go dependencies
FROM golang:latest as builder

ADD . /chain-monitor
ENV GOPROXY="https://goproxy.cn,direct"
RUN cd /chain-monitor && \
    go mod tidy &&  \
    go build -v -o /chain-monitor/build/bin/chain-monitor ./cmd/

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

COPY --from=builder /chain-monitor/build/bin/chain-monitor /bin/

ENTRYPOINT ["chain-monitor"]