FROM golang:1.20 AS builder

WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Inisialisasi module dan download dependency
RUN go mod tidy

# Build aplikasi dengan path yang benar
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd  # Tambahkan CGO_ENABLED=0

# Gunakan Alpine sebagai base image untuk aplikasi yang lebih kecil
FROM alpine
WORKDIR /root/
COPY --from=builder /app/main .

# Beri izin eksekusi pada binary
RUN chmod +x main

# Expose port API
EXPOSE ${API_PORT}
CMD ["./main"]
