# --- Build stage ---
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /legaltech-server ./cmd/server

# --- Run stage ---
FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /legaltech-server .
COPY firebase-service-account.json ./
EXPOSE 8080
CMD ["./legaltech-server"]