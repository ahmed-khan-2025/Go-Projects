FROM golang:1.26-alpine

WORKDIR /app

# Copy go mod files from ROOT context
COPY go.mod go.sum ./
RUN go mod download

# Copy full project
COPY . .

# Build only the order-service
RUN go build -o main ./services/order-service

CMD ["./main"]