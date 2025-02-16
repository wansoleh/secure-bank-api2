FROM golang:1.18 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

FROM alpine
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE ${API_PORT}
CMD ["./main"]
