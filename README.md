# 🚀 gRPC Order System

A scalable microservices-based system built with **Go** and **gRPC**, simulating a real-world distributed architecture for order processing.

---

## 🧠 Overview

This project demonstrates how to design and implement a **distributed system** using:

* Go
* gRPC
* Protocol Buffers
* Clean Architecture principles

The system is composed of multiple services communicating via **high-performance RPC**.

---

## 🏗️ Architecture

```
                +------------------+
                |   API Gateway    |
                |  (optional REST) |
                +--------+---------+
                         |
                         v
                +------------------+
                |  Order Service   |
                +--------+---------+
                         |
        +----------------+----------------+
        |                                 |
        v                                 v
+---------------+                +------------------+
| User Service  |                | Inventory Service|
+---------------+                +------------------+
```

---

## 📦 Project Structure

```
grpc-order-system/
├── proto/
├── services/
│   ├── order-service/
│   ├── user-service/
│   └── inventory-service/
├── pkg/
├── docker/
└── README.md
```

---

## ⚙️ Tech Stack

* **Language:** Go
* **RPC:** gRPC
* **Serialization:** Protocol Buffers
* **Database:** PostgreSQL (planned)
* **Cache:** Redis (planned)
* **Containerization:** Docker

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/seuuser/grpc-order-system.git
cd grpc-order-system
```

---

### 2. Install dependencies

```bash
go mod tidy
```

---

### 3. Install Protocol Buffers

```bash
sudo apt install -y protobuf-compiler
```

---

### 4. Install gRPC plugins

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

---

### 5. Generate gRPC code

```bash
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

---

### 6. Run the Order Service

```bash
go run services/order-service/cmd/main.go
```

---

## 🧪 Testing the service

You can use **grpcurl** to test:

```bash
grpcurl -plaintext localhost:50051 list
```

---

## 🔥 Features

* gRPC-based communication between services
* Modular and scalable architecture
* Separation of concerns (handler, service, repository)
* Ready for microservices evolution

---

## 🧩 Roadmap

* [ ] Add User Service
* [ ] Add Inventory Service
* [ ] Implement database integration
* [ ] Add authentication & authorization
* [ ] Implement gRPC interceptors
* [ ] Add observability (logs, metrics, tracing)
* [ ] Docker Compose environment
* [ ] CI/CD pipeline

---

## 🔐 Future Improvements

* TLS security for gRPC
* Circuit breaker & retries
* API Gateway with REST support
* Integration with message brokers

---

## 🤝 Contributing

Feel free to open issues or submit pull requests.

---

## 👩‍💻 Author

**Ingrid Caroline**
📧 [ingridcaroline725@gmail.com](mailto:ingridcaroline725@gmail.com)
🌎 Brazil

---

## ⭐ Why this project?

This project was designed to showcase:

* Backend engineering skills
* Distributed systems knowledge
* Production-ready architecture patterns

---
