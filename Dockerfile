# Build stage
FROM golang:1.24.4-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server main.go

# Production stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server ./
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/assets ./assets
EXPOSE 8080
CMD ["./server"]