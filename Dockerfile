# Etapa de build
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /poc-go-eks ./cmd/main.go

# Etapa final
FROM alpine:3.17
WORKDIR /app

COPY --from=builder /poc-go-eks /app/poc-go-eks

EXPOSE 8080
ENTRYPOINT ["/app/poc-go-eks"]
