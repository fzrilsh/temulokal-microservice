# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY . .
# Optional: sync workspace if available (no-op if not)
RUN go work sync || true
WORKDIR /src/auth-service
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/auth-service .

# Runtime stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /out/auth-service /app/auth-service
EXPOSE 8001
# Run the service with cobra command "start"
ENTRYPOINT ["/app/auth-service","start"]
