FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .
RUN  apk add --no-cache libpcap-dev && go mod tidy && go build -o arp-spoofing-go .


FROM alpine:3.14 

COPY --from=builder /build/arp-spoofing-go /app
COPY --from=builder /build/conf   /app/conf
COPY --from=builder /build/logo/logo.txt /root/src/arp-spoofing/logo/logo.txt
ENTRYPOINT [ "/app/arp-spoofing-go" ]
CMD []