# 🐹 Personality Manager API

Welcome to the **Personality Manager API** — a simple and elegant REST API for managing personality data, built with **Golang**, **Mux**, and **PostgreSQL**.

## 📖 Features
- 🚀 Lightweight and fast API powered by Go's performance.
- 🛠️ Seamlessly handle CRUD operations for personalities.
- 🔗 Integrated with PostgreSQL for robust data persistence.
- 🐳 Ready-to-use Docker Compose configuration.

---

![Personalities API](https://preview.redd.it/tcmyd3n69ng41.jpg?width=1080&crop=smart&auto=webp&s=842708a5e9c47fb8ea6c2c639a72dc131691891b)

---

## ✅ Prerequisites
Make sure you have the following installed:
1. **Docker** & **Docker Compose**: For containerized setup.
2. **Go** (1.19+): If you wish to run without Docker.

---

## 🛠 Getting Started

### 🛠 Built With
- **Go**: Programming language.
- **Mux**: HTTP router.
- **PostgreSQL**: Database.
- **Docker Compose**: Containerization.

### Clone the Repository
```bash
git clone https://github.com/rdiego26/golang-simple-rest-api.git
cd golang-simple-rest-api
```

### Run with Docker Compose (database)
```bash
docker-compose up
```

### Run Application
```bash
go run main.go
```

### API Endpoints
- **GET** `/`: Home route.
- **GET** `/api/personalities`: Get all personalities.
- **GET** `/api/personalities/{id}`: Get a personality by ID.
- **POST** `/api/personalities`: Create a new personality.
- **PUT** `/api/personalities/{id}`: Update a personality by ID.
- **DELETE** `/api/personalities/{id}`: Delete a personality by ID.
