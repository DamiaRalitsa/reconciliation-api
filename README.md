# reconciliation-api
simple reconciliation/ service

📦 Dependencies
Go 1.22+

Fiber Web Framework

Standard Library only

- REST API built using [Fiber](https://gofiber.io/)
- Docker-ready & config-driven

## 📂 Project Structure
No worries, Damia — here’s a clean and beautiful `Project Structure` block you can paste into your `README.md` so it looks neat on GitHub:

---

```markdown
## 📂 Project Structure

```
reconciliation-api/
├── assets/
│   └── templates/                      # Sample CSV templates
│       ├── bank\_bca\_statements.csv
│       ├── bank\_mandiri\_statements.csv
│       └── system\_transaction.csv
│
├── cmd/
│   └── app/
│       └── main.go                     # App entrypoint
│
├── config.json                         # App configuration
├── Dockerfile                          # Docker build config
├── Makefile                            # Optional build script
│
├── internal/
│   ├── delivery/http/                  # HTTP delivery layer
│   │   ├── route/                      # API route config
│   │   └── reconciliation\_controller.go
│   ├── domain/                         # Domain models & enums
│   ├── repositories/reconciliation/    # CSV parsing
│   ├── usecases/reconciliation/        # Reconciliation logic
│   └── helper.go                       # Shared helpers
│
├── docs/
│   └── Simple Reconciliation Service.postman\_collection.json
│
├── go.mod
├── go.sum
├── README.md
```

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

