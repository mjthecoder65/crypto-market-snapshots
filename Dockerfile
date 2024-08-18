# Build stage
FROM golang:1.23-alpine3.20 AS builder

WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/main .

CMD [ "/app/main" ]