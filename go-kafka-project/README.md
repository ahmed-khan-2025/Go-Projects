# Go Kafka Project

## 📌 Overview

This project demonstrates a basic integration of Apache Kafka with Go. It includes a simple producer and consumer setup to send and receive messages using Kafka.

## 🚀 Features

* Kafka producer in Go
* Kafka consumer in Go
* JSON-based message structure
* Simple event-driven communication

## 🛠️ Tech Stack

* Go (Golang)
* Apache Kafka
* Docker (optional)

## 📂 Project Structure

```
go-kafka-project/
├── producer/
├── consumer/
├── shared/
└── main.go
```

## ⚙️ Setup & Run

### 1. Start Kafka (Docker)

```bash
docker compose up -d
```

### 2. Create Topic

```bash
docker exec -it kafka kafka-topics --create --topic events --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
```

### 3. Run Producer

```bash
go run producer/main.go
```

### 4. Run Consumer

```bash
go run consumer/main.go
```

## 📊 Example Event

```json
{
  "id": "1",
  "type": "sensor",
  "value": "temperature",
  "time": "2026-05-01T22:30:04+02:00"
}
```

## 🎯 Purpose

This project serves as a **learning foundation** for understanding:

* Kafka basics
* Producer/Consumer model
* Event-driven communication

## 📌 Next Step

See the advanced version: `go-kafka-microservices`

