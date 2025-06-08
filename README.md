# Coffee Shop API

A RESTful API for managing coffee products, built with Go, Gin, and Swagger.

## Features

- CRUD operations for coffee products
- Swagger documentation
- Clean architecture
- Dependency injection with Google Wire
- In-memory data store (easily replaceable with a real database)

## Prerequisites

- Go 1.22 or later
- Swag CLI for Swagger documentation generation

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd coffee-shop-api
```

2. Install dependencies:
```bash
go mod download
```

3. Install Swag CLI (if not already installed):
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## Generating Swagger Documentation

To generate the Swagger documentation, run:
```bash
make swagger
```

This will generate the documentation in the correct location (`internal/docs/`).

## Running the Application

To start the server:
```bash
make run
```

The server will start on `http://localhost:8080`

## API Documentation

Once the server is running, you can access the Swagger UI at:
```
http://localhost:8080/swagger/index.html
```

## API Endpoints

- `GET /coffees` - List all coffees
- `GET /coffees/:id` - Get a specific coffee
- `POST /coffees` - Create a new coffee
- `PUT /coffees/:id` - Update a coffee
- `DELETE /coffees/:id` - Delete a coffee

## Project Structure

```
coffee-shop-api/
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── handler/          # HTTP handlers
│   ├── service/          # Business logic
│   ├── repository/       # Data access
│   ├── model/           # Data models
│   ├── di/              # Dependency injection
│   └── docs/            # Swagger documentation
├── go.mod
└── README.md
```

## License

MIT 

## Development

For the best development experience, use the development workflow:

```bash
make dev
```

This single command will:
1. Install required development tools (air, swag)
2. Install project dependencies
3. Generate Swagger documentation
4. Start the development server with hot-reload

The server will automatically restart when you make changes to your code.

### Debug Mode

To run the application in debug mode:

```bash
make debug
```

This will:
1. Start the application with Delve debugger
2. Listen on port 2345 for debugger connections
3. Allow you to connect with your IDE's debugger

To connect with VS Code:
1. Install the Go extension
2. Create a launch configuration in `.vscode/launch.json`:
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Debug Coffee Shop API",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "remotePath": "",
            "port": 2345,
            "host": "127.0.0.1",
            "showLog": true
        }
    ]
}
```

## Available Commands

The project uses a Makefile to simplify common tasks:

- `make dev` - Start development server with hot-reload (recommended for development)
- `make debug` - Start development server in debug mode
- `make deps` - Install dependencies
- `make build` - Build the application
- `make run` - Run the application
- `make clean` - Clean build files
- `make test` - Run tests
- `make swagger` - Generate Swagger documentation
- `make all` - Clean and build the application
- `make help` - Show available commands 