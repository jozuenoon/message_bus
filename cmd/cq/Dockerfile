FROM golang:1.12-alpine as builder
COPY . /build
WORKDIR /build

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN apk add git coreutils wget \
    && go mod vendor
RUN go build -o bin/server cmd/cq/main.go

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.0 && \
    wget -qO/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /grpc_health_probe

FROM scratch
ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /build/bin/server ./
COPY --from=builder /grpc_health_probe ./

EXPOSE 8000 9000
CMD ["/server"]