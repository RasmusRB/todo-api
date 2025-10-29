# Todo API

A simple RESTful Todo API built with Go and Gin framework, featuring Swagger documentation and Docker support.

## Features

- ✅ RESTful API endpoints for Todo CRUD operations
- 📚 Swagger/OpenAPI documentation
- 🌐 CORS middleware
- 🐳 Docker containerization
- 🏗️ Clean architecture with separate packages

## Prerequisites

- Go 1.25.3 or higher
- Docker (optional, for containerization)

## Project Structure

```
gotut/
├── build/
│   └── Dockerfile          # Docker configuration
├── docs/                   # Auto-generated Swagger docs
├── handlers/               # API handlers
│   └── handlers.go
├── middleware/             # Custom middleware
│   └── middleware.go
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
└── README.md
```

## Installation

### Local Development

1. Clone the repository:

```bash
git clone <repository-url>
cd gotut
```

2. Install dependencies:

```bash
go mod download
```

3. Run the application:

```bash
go run .
```

The server will start on `http://localhost:8080`

## API Endpoints

| Method | Endpoint              | Description              |
| ------ | --------------------- | ------------------------ |
| GET    | `/todos`              | Get all todos            |
| POST   | `/todos`              | Create a new todo        |
| PUT    | `/todos/:id`          | Update a todo by ID      |
| DELETE | `/todos/:id`          | Delete a todo by ID      |
| GET    | `/swagger/index.html` | Swagger UI documentation |

## API Documentation

Once the server is running, access the interactive Swagger documentation at:

```
http://localhost:8080/swagger/index.html
```

### Regenerate Swagger Documentation

If you modify the API handlers, regenerate the Swagger docs:

```bash
swag init
```

## Docker

### Build Docker Image

```bash
docker build -t gotut-api -f build/Dockerfile .
```

### Run Docker Container

```bash
docker run -p 8080:8080 gotut-api
```

## Example Usage

### Get All Todos

```bash
curl http://localhost:8080/todos
```

### Create a Todo

```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"id":"4","title":"Learn Go","done":false,"detail":"Complete Go tutorial"}'
```

### Update a Todo

```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"id":"1","title":"Buy groceries","done":true,"detail":"Completed shopping"}'
```

### Delete a Todo

```bash
curl -X DELETE http://localhost:8080/todos/1
```

## Development

### Project Packages

- **handlers**: Contains all HTTP request handlers
- **middleware**: Custom middleware (CORS, etc.)
- **docs**: Auto-generated Swagger documentation

### Adding New Endpoints

1. Add handler function in `handlers/handlers.go`
2. Add Swagger annotations above the handler
3. Register route in `main.go`
4. Run `swag init` to update documentation

## Technologies Used

- [Go](https://golang.org/) - Programming language
- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [Swaggo](https://github.com/swaggo/swag) - Swagger documentation generator
- [Docker](https://www.docker.com/) - Containerization

## License

This project is open source and available under the [MIT License](LICENSE).
