# Build stage
FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -a -installsuffix cgo -o main ./cmd/main.go

# Production stage
FROM alpine:3.17

WORKDIR /root/

COPY --from=builder /app/main .
COPY .env .env

EXPOSE 8080

CMD ["./main"]