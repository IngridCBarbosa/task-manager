# 📦 Task Manager API (Golang)

This is a simple Task Manager API built with Go, created with the primary goal of studying and applying the **Flat Folder Structure** architectural pattern.

The project also uses Redis as the database layer to better understand data persistence, indexing strategies, and communication with external services in a lightweight architecture.

---

## 🎯 Purpose

The focus of this project is **not business complexity**, but code organization and architectural learning.

The main idea is to understand in practice:

- How to structure a project without complex layered architecture
- When a flat structure is sufficient
- The limitations of this approach as the system grows
- How Redis can be used as a lightweight data store

---

## 🧱 Architecture: Flat Folder Structure

In this project, all files are organized at the same level, avoiding deep hierarchies such as `internal/`, `domain/`, or `usecase/`.

This provides a simpler and more direct view of the application.

```bash
task-manager/
├── main.go
├── handler.go
├── service.go
├── repository.go
├── model.go
├── db.go
├── routes.go
├── Makefile
└── go.mod