# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY . .
RUN go work sync || true
WORKDIR /src/gateway
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/gateway .

# Runtime stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /out/gateway /app/gateway
EXPOSE 8000
ENTRYPOINT ["/app/gateway","start"]
