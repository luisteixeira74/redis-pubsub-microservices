# Stage 1: build
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Stage 2: final image (alpine mais leve)
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

ENTRYPOINT ["./app"]
