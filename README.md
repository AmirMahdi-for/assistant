# Assistant - Go Migration (Step-by-Step)

This project is a step-by-step migration of the Shopping Assistant backend from Laravel to Go, starting with the message storing feature using the MVC pattern and the Gin framework.

---

## 🛠️ Prerequisites

- **Go** version `go1.23.4`
  Download and install from [https://golang.org/dl](https://golang.org/dl)

---

## 🚀 Getting Started

### 1. Initialize the Project and Install Gin

Inside your project folder, run the following:

```bash
go mod init assistant
go get -u github.com/gin-gonic/gin
go get github.com/joho/godotenv

