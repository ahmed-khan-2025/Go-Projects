# Go Kafka Microservices 🚀

A scalable event-driven microservices architecture built with Go (Golang) and Apache Kafka.
This project demonstrates real-world distributed system design using asynchronous messaging between services.

---

## 📌 Overview

This system simulates a real-time event-driven architecture where multiple microservices communicate via Kafka topics.

Each service is responsible for a specific domain and processes events independently.

Built using **Go + Apache Kafka + Docker**.

---

## 🧠 Architecture

```text id="arch_01"
Order Service → Kafka → Analytics Service
Sensor Service → Kafka → Storage Service
                → Kafka → Notification Service
```

All services communicate asynchronously using Kafka events.

---

## ✨ Microservices

### 🛒 Order Service

* Produces order-related events
* Sends events to Kafka

### 📡 Sensor Service

* Simulates IoT/sensor data
* Streams real-time events

### 📊 Analytics Service

* Consumes events
* Processes and analyzes data

### 🔔 Notification Service

* Listens to events
* Sends notifications (simulated)

### 💾 Storage Service

* Persists event data
* Acts as a data sink

---

## 🛠️ Tech Stack

* Go (Golang)
* Apache Kafka
* Docker & Docker Compose
* Event-driven architecture

---

## 📂 Project Structure

```text id="structure_01"
go-kafka-microservices/
│
├── docker-compose.yml
├── .env
│
├── services/
│   ├── order-service/
│   │   └── main.go
│   ├── sensor-service/
│   │   └── main.go
│   ├── analytics-service/
│   │   └── main.go
│   ├── notification-service/
│   │   └── main.go
│   └── storage-service/
│       └── main.go
│
├── internal/
│   ├── kafka/
│   │   └── kafka.go
│   └── models/
│       └── event.go
```

---

## 🚀 How to Run

### 1. Clone repository

```bash id="clone_01"
git clone https://github.com/<your-username>/go-kafka-microservices.git
cd go-kafka-microservices
```

---

### 2. Start Kafka + services

```bash id="run_01"
docker compose up --build
```

---

### 3. Check logs

Each service will start producing/consuming events automatically.

---

## 🔄 Event Flow

1. Services generate events (Order, Sensor)
2. Events are published to Kafka topics
3. Consumer services process events
4. Results are stored, analyzed, or used for notifications

---

## 📦 Kafka Topics

| Topic         | Purpose          |
| ------------- | ---------------- |
| orders        | Order events     |
| sensors       | Sensor data      |
| analytics     | Processed data   |
| notifications | Alerts/messages  |
| storage       | Persisted events |

---

## 🧠 Key Concepts Demonstrated

* Microservices architecture
* Event-driven design
* Asynchronous communication
* Kafka pub/sub model
* Decoupled system design
* Scalable backend architecture

---

## 📌 Learning Outcomes

This project demonstrates:

* Building distributed systems in Go
* Working with Kafka event streams
* Designing scalable microservices
* Handling asynchronous workflows

---

## 🔥 Future Improvements

* Add database persistence (PostgreSQL / MongoDB)
* Add API Gateway
* Add authentication (JWT)
* Add monitoring (Prometheus + Grafana)
* Deploy using Kubernetes

---

## ⭐ Purpose

This project is part of a backend engineering learning path focused on distributed systems, microservices, and event-driven architecture.

