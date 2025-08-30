# 🧾 Inventory Management System – Backend API (Go + Gin)

A backend-only inventory management system built with Go and the Gin framework. It provides secure RESTful APIs for managing products, orders, and user authentication. Designed for scalability, modularity, and integration with any frontend or third-party system.

---

## 🚀 Features

### 🔐 Authentication
- Role-based access using middleware (`admin`, `member`)
- JWT or token-based authentication (assumed via `AuthMiddleware`)

---

### 📦 Product Management (`member` access)
| Method | Endpoint                     | Description                          |
|--------|------------------------------|--------------------------------------|
| POST   | `/Products`                  | Add a new product                    |
| GET    | `/Products/:id`              | View product by ID                   |
| GET    | `/Products`                  | View all products                    |
| PUT    | `/Products/:id`              | Update product details               |
| DELETE | `/Products/:id`              | Delete a product                     |
| DELETE | `/Instock/`                  | Mark product as in-stock             |
| POST   | `/Inuse`                     | Mark product as in-use               |
| GET    | `/Username/:use_by`          | Filter products by user              |
| GET    | `/inventory_view?status=inuse` | View inventory by status           |

---

### 📑 Order Management (`admin` access)
| Method | Endpoint         | Description               |
|--------|------------------|---------------------------|
| POST   | `/orders`        | Create a new order        |
| GET    | `/orders`        | Get all orders            |
| GET    | `/orders/:id`    | Get order by ID           |
| PUT    | `/orders/:id`    | Update order              |
| DELETE | `/orders/:id`    | Delete order              |

---

### 👤 User Authentication
| Method | Endpoint   | Description        |
|--------|------------|--------------------|
| POST   | `/Signup`  | Register new user  |
| POST   | `/Login`   | Authenticate user  |

---

## 🧱 Tech Stack

- **Language**: Go (Golang)
- **Framework**: Gin
- **Database**: MySQL
- **Architecture**: Modular (Controllers, Middleware, Routing)
- **Concurrency**: Goroutines, WaitGroups, Worker Pools (used in bulk operations)

---

## 📁 Project Structure (Simplified)

