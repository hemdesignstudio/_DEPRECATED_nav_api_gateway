FROM golang:1.11-alpine as dependencies

RUN apk --no-cache add curl git gcc musl-dev

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN mkdir -p /go/src/github.com/hem-nav-gateway
ADD . /go/src/github.com/hem-nav-gateway/
WORKDIR /go/src/github.com/hem-nav-gateway

RUN dep ensure -vendor-only

# Tester
FROM dependencies as tester
RUN go test ./... -v

# Builder
FROM dependencies as builder
RUN go build -o main .

# Release
FROM alpine

RUN apk --no-cache --update add ca-certificates

COPY --from=builder /go/src/github.com/hem-nav-gateway/main /
COPY --from=builder /go/src/github.com/hem-nav-gateway/cert/rootCA.pem /cert/rootCA.pem

EXPOSE 6789

CMD ["./main"]

