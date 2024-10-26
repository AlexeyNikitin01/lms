FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o user ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/user .
COPY etc/config.docker.yml etc/config.template.yml
EXPOSE 18080 50054

CMD ["./user"]
