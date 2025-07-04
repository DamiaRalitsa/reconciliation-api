# reconciliation/-api
simple reconciliation/ service

📦 Dependencies
Go 1.22+

Fiber Web Framework

Standard Library only

- REST API built using [Fiber](https://gofiber.io/)
- Docker-ready & config-driven

## 📂 Project Structure
├── assets/templates/ # Sample CSV templates (system, BCA, Mandiri)
│ ├── system_transaction.csv
│ ├── bank_bca_statements.csv
│ └── bank_mandiri_statements.csv
│
├── cmd/app/main.go # App entrypoint
│
├── config.json # App configuration
├── Dockerfile # Docker build config
├── Makefile # Optional build helper
│
├── internal/
│ ├── delivery/http/ # HTTP controllers
│ │ └── route/ # API route definitions
│ ├── domain/ # Domain models & types
│ ├── repositories/reconciliation/ # CSV parsing & data loading
│ └── usecases/reconciliation/ # Reconciliation use case logic
│
├── helper.go # Shared helpers
├── README.md # Project docs
├── docs/
│ └── Simple Reconciliation Service.postman_collection.json
│
├── go.mod
├── go.sum

---

## 🛠️ Running Locally

### 1. Run with Go:

```bash
go run cmd/app/main.go
```
```bash
docker build -t reconciliation-api . && docker image prune -f && docker run -it --rm   --network app-network   -p 8444:8444   -v $(pwd)/config.json:/config.json   reconciliation-api
```

👩‍💻 Author
Damia Ralitsa
Software Engineer / Golang Developer

