FROM golang:1.26-alpine

WORKDIR /app

# Copy dependency files from ROOT context
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project
COPY . .

# Build ONLY storage-service
RUN go build -o main ./services/storage-service

CMD ["./main"]