FROM golang:1.26-alpine

WORKDIR /app

# Copy dependency files from ROOT context
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project
COPY . .

# Build only analytics service
RUN go build -o main ./services/analytics-service

CMD ["./main"]