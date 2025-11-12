FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

FROM alpine:3.18

RUN apk add --no-cache ca-certificates bash netcat-openbsd

WORKDIR /app

COPY --from=builder /app/main .

COPY .env .
#COPY consumer.kafka.conf.yml .

RUN mkdir logs/

EXPOSE 8080

# CMD ["./bin/wait-for-it.sh", "localhost", "9092", "./main"]
CMD ["./main"]
