# Stage 1: build
FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app .

# Stage 2: final image
FROM golang:1.22

WORKDIR /app

COPY --from=builder /app/app .

ENTRYPOINT ["./app"]
