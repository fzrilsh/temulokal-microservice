# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY . .
RUN go work sync || true
WORKDIR /src/umkm-service
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/umkm-service .

# Runtime stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /out/umkm-service /app/umkm-service
EXPOSE 8002
ENTRYPOINT ["/app/umkm-service","start"]
