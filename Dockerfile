# Stage 1: Build stage
FROM golang:1.23-alpine AS builder

# Set working directory di dalam container
WORKDIR /app

# Copy semua file aplikasi ke dalam container
COPY . .
# Copy go mod and sum files
# COPY go.mod go.sum ./

# Download dan install dependensi Go
RUN go mod tidy

# Build aplikasi Go
RUN go build -o backendApp .

# Stage 2: Final stage
FROM alpine:latest

# Set working directory di dalam container
WORKDIR /root/

# Copy file binary hasil build dari stage 1 ke image final
COPY --from=builder /app/backendApp .
# Menyalin file .env ke dalam container
COPY .env .env  

# Expose port yang akan digunakan oleh aplikasi (misalnya port 8088)
EXPOSE 80

# Perintah untuk menjalankan aplikasi Go
CMD ["./backendApp"]
