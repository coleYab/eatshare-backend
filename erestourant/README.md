# 🥘 eRestourant

A simple restaurant backend service built with [Go](https://golang.org/) and the [Gin](https://github.com/gin-gonic/gin) web framework. It supports environment-based configuration and PostgreSQL integration.

---

## 📁 Project Structure

```
.
├── build/              # Compiled binary output
├── cmd/
│   └── api/            # Main application entry point
│       └── main.go
├── config/             # Configuration-related code
├── internal/           # Application logic (handlers, services, etc.)
├── sqlc.yaml           # SQLC config for database interaction
├── go.mod / go.sum     # Go module definitions
├── .env                # Environment variables
├── makefile            # Build/run automation
└── README.md           # Project documentation
```

---

## 🚀 Getting Started

### Prerequisites

* [Go](https://golang.org/doc/install) 1.18+
* [PostgreSQL](https://www.postgresql.org/)
* [sqlc](https://docs.sqlc.dev/en/stable/) (optional, if using codegen for queries)

---

## ⚙️ Configuration

Create a `.env` file in the root directory:

```env
PORT=5000
ENV="DEVELOPMENT"
DB_CONN_URI="postgres://postgres:postgres@localhost:5432/erestourant"
```

---

## 🛠️ Build & Run

You can build and run the project using the `makefile`:

### Build

```bash
make build
```

### Run

```bash
make run
```

This will compile the code and run the server at:

```
http://localhost:5000
```

---

## 🧹 Clean Build Artifacts

```bash
make clear
```

---

## 🗃️ Database

The project expects a PostgreSQL database running with the following default config (can be changed in `.env`):

* **User:** `postgres`
* **Password:** `postgres`
* **Database:** `erestourant`
* **Host:** `localhost`
* **Port:** `5432`

---

## 📦 Dependencies

* [Gin](https://github.com/gin-gonic/gin)

Install dependencies via:

```bash
go mod tidy
```

---

## 📌 Routes Overview

The API is organized into modular route groups:

### 🔐 Auth Routes

> Handles user authentication

| Method | Endpoint         | Description         |
| ------ | ---------------- | ------------------- |
| POST   | `/auth/login`    | Log in a user       |
| POST   | `/auth/register` | Register a new user |

---

### 📋 Menu Routes

> Manages menu items in the restaurant

| Method | Endpoint    | Description                      |
| ------ | ----------- | -------------------------------- |
| GET    | `/menu/`    | Get all menu items               |
| GET    | `/menu/:id` | Get details of a menu item by ID |
| POST   | `/menu/`    | Create a new menu item           |
| PUT    | `/menu/:id` | Update a menu item by ID         |
| DELETE | `/menu/:id` | Delete a menu item by ID         |

---

### 🛒 Order Routes

> Manages customer orders

| Method | Endpoint            | Description              |
| ------ | ------------------- | ------------------------ |
| GET    | `/order/`           | Get all orders           |
| GET    | `/order/:id`        | Get order details by ID  |
| GET    | `/checkout/order/:id` | Marks an order a complete |
| POST   | `/order/`           | Create a new order       |
| PUT    | `/order/:id`        | Update an existing order |
| DELETE | `/order/:id`        | Delete an order by ID    |

---

### 👤 User Routes

> Handles user management

| Method | Endpoint    | Description            |
| ------ | ----------- | ---------------------- |
| GET    | `/user/`    | Get all users          |
| GET    | `/user/:id` | Get user details by ID |
| DELETE | `/user/:id` | Delete a user by ID    |

---

## 📌 TODO
* Implement authentication middeware
* Implement some admin features

