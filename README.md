# ProductAPI

ProductAPI is an API for users to create, read, update, and delete products, featuring a backend API written in Go and a frontend using JavaScript, CSS, and HTML. The project demonstrates full-stack capabilities, including database connectivity, authentication, and a modular architecture.

## Table of Contents

- [Project Structure](#project-structure)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
- [Environment Variables](#environment-variables)
- [Scripts & Usage](#scripts--usage)
- [Contributing](#contributing)
- [License](#license)

## Project Structure

```
.
├── .env
├── cmd/           # Entrypoint commands for the application
├── docs/          # Documentation files
├── frontend/      # Frontend code (JavaScript, CSS, HTML)
├── go.mod
├── go.sum
├── internal/      # Go internal packages (business logic, API handlers, etc.)
└── sqlc.yaml      # SQLC config for Go database code generation
```

## Tech Stack

- **Go** (Backend, API, business logic) – 60.2%
- **JavaScript** (Frontend logic) – 33.9%
- **CSS** (Styling) – 4.9%
- **HTML** (Markup) – 1%

## Getting Started

### Prerequisites

- Go 1.18+
- Node.js & npm (for frontend dependencies)
- PostgreSQL (or compatible database)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/Black-tag/ProductAPI.git
   cd ProductAPI
   ```

2. Copy and configure environment variables:
   ```sh
   cp .env.example .env
   # Edit .env as needed
   ```

3. Install Go dependencies:
   ```sh
   go mod download
   ```

4. Setup and run the database (PostgreSQL recommended).

5. (Optional) Install frontend dependencies:
   ```sh
   cd frontend
   npm install
   ```

### Running the Application

- **Backend:**
  ```sh
  go run cmd/main.go
  ```
- **Frontend:**
  ```sh
  cd frontend
  npm start
  ```

## Environment Variables

Copy `.env.example` to `.env` and fill in your configuration. Example:

```env
DB_URL="postgresql://dbusername:dbpassword@localhost:5432/dbname?sslmode=disable"
SECRET="W2X3k961oS5JUYMoyHk+BWnXKQ4u21rSH8eSPpnw9mdkp0+Xq4iX5oprtMFjF9Zp
3OnKhfbhCxAiM9vCKoPa/A=="
```

| Variable | Description                                 | Example Value                                                |
|----------|---------------------------------------------|--------------------------------------------------------------|
| DB_URL   | Database connection string (PostgreSQL)     | postgresql://dbusername:dbpassword@localhost:5432/dbname...  |
| SECRET   | Secret key for authentication/session/token | W2X3k961oS5JUYMoyHk+...                                      |



## Scripts & Usage

- **Start backend:** `go run cmd/main.go`
- **Start frontend:** `npm start` (inside `frontend/`)


## Contributing

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/foo`)
3. Commit your changes
4. Push to the branch (`git push origin feature/foo`)
5. Open a Pull Request



