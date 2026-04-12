FROM golang:1.26-alpine AS builder

WORKDIR /src
COPY protos /protos
COPY sso/go.mod sso/go.sum ./
RUN go mod download

COPY sso ./
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /sso ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /sso /app/
COPY sso/config/local.yaml /app/config/config.yaml

EXPOSE 50051
ENTRYPOINT ["/app/sso"]
