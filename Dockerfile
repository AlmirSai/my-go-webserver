# Используем легковесный образ Go
FROM golang:1.23.3-alpine as builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o webserver ./cmd/webserver/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/webserver .

EXPOSE 8080

CMD ["./webserver"]
