<img src="https://go.dev/images/go-logo-blue.svg" width="150" height="150">

[![Go](https://github.com/LeoCourbassier/golang-grpc-template/actions/workflows/go.yml/badge.svg)](https://github.com/LeoCourbassier/golang-grpc-template/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/LeoCourbassier/golang-grpc-template)](https://goreportcard.com/report/github.com/LeoCourbassier/golang-grpc-template)
[![codecov](https://codecov.io/gh/LeoCourbassier/golang-grpc-template/branch/main/graph/badge.svg?token=3Z3Z3Z3Z3Z)](https://codecov.io/gh/LeoCourbassier/golang-grpc-template)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# Go gRPC Template

> [!NOTE]
> This project is a work in progress. Some features are not yet implemented.
>
> The gRPC clients do not use credentials by default. This is not secure for production use.
>
> Contributions are welcome!

A template for building gRPC modular monoliths in Go. Designed as a starting point for new projects with all boilerplate already configured.

## Project Structure

```
├── internal/                # Private packages (not importable by other projects)
│   ├── clients/             # gRPC clients for inter-module communication
│   ├── config/              # Configuration management (Viper)
│   ├── constants/           # Application constants
│   ├── controller/          # Controller registration utilities
│   ├── db/                  # Database connection and initialization
│   ├── errors/              # Custom error definitions
│   ├── logger/              # Logging setup
│   │   └── formatters/      # Log formatters (JSON, pretty, text)
│   └── mocks/               # Generated mocks for testing
│
├── pkg/                     # Public packages (main application modules)
│   ├── ping/                # Health check module
│   │   ├── controllers/     # gRPC service handlers
│   │   └── proto/           # Generated protobuf code
│   ├── users/               # Users domain module
│   │   ├── application/     # Business logic (service layer)
│   │   ├── controllers/     # gRPC service handlers
│   │   ├── models/          # Domain models and interfaces
│   │   ├── repository/      # Data access layer
│   │   ├── proto/           # Generated protobuf code
│   │   └── views/           # Request/response DTOs
│   └── server/              # gRPC server setup and registration
│
├── docker/                  # Docker-related files
│   └── db/                  # Database initialization scripts
│
├── main.go                  # Application entry point
├── config.yaml              # Runtime configuration
├── Taskfile.yaml            # Task runner automation
├── Dockerfile               # Multi-stage Docker build
└── docker-compose.yml       # Docker Compose configuration
```

### Module Structure

Each module in `pkg/` follows a layered architecture:

| Layer | Purpose |
|-------|---------|
| `controllers/` | gRPC service handlers |
| `application/` | Business logic and orchestration |
| `repository/` | Data access and persistence |
| `models/` | Domain entities and interfaces |
| `views/` | Request/response DTOs |
| `proto/` | Generated protobuf code |

## Prerequisites

- [Go 1.24+](https://golang.org/dl/)
- [Task](https://taskfile.dev/) (task runner)
- [Docker](https://www.docker.com/) (for database)
- [protoc](https://grpc.io/docs/protoc-installation/) (protocol buffer compiler)

## Getting Started

1. Clone the repository
   ```bash
   git clone https://github.com/LeoCourbassier/golang-grpc-template.git
   cd golang-grpc-template
   ```

2. Install dependencies
   ```bash
   go mod tidy
   ```

3. Start the database
   ```bash
   task db
   ```

4. Run the server
   ```bash
   task run
   ```

## Available Tasks

| Command | Description |
|---------|-------------|
| `task run` | Generate mocks and start the gRPC server |
| `task build` | Generate mocks and build the binary |
| `task test` | Run tests with coverage report |
| `task gen` | Generate mocks using mockgen |
| `task db` | Start PostgreSQL via docker-compose |

## Testing

Run all tests with coverage:
```bash
task test
```

This generates a coverage report and opens it in your browser.

### Mock Generation

Mocks are generated using [mockgen](https://github.com/uber-go/mock):
```bash
task gen
```

Generated mocks are placed in `internal/mocks/`.

## Docker

### Start Database Only
```bash
task db
# or
docker-compose up -d db
```

### Start Full Application
```bash
docker-compose up --build
```

The gRPC server runs on port `50051` and PostgreSQL on port `5432`.

## Configuration

Configuration is managed via `config.yaml`:

```yaml
database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "password"
  name: "postgres"

logger:
  level: "info"        # debug, info, warn, error
  format: "pretty"     # pretty, json, text

server:
  port: 50051
  host: "localhost"
  clients:             # Inter-module gRPC clients
    - name: "ping"
      host: "localhost"
      port: 50051
    - name: "users"
      host: "localhost"
      port: 50051
```

## Dependencies

| Package | Purpose |
|---------|---------|
| [google.golang.org/grpc](https://pkg.go.dev/google.golang.org/grpc) | gRPC framework |
| [google.golang.org/protobuf](https://pkg.go.dev/google.golang.org/protobuf) | Protocol buffers |
| [gorm.io/gorm](https://gorm.io/) | ORM for database operations |
| [go.uber.org/fx](https://pkg.go.dev/go.uber.org/fx) | Dependency injection |
| [github.com/spf13/viper](https://github.com/spf13/viper) | Configuration management |
| [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) | Structured logging |
| [github.com/go-playground/validator](https://github.com/go-playground/validator) | Input validation |
| [go.uber.org/mock](https://github.com/uber-go/mock) | Mock generation for testing |

## Features

- **Modular Monolith**: Self-contained modules in `pkg/` with clear boundaries
- **gRPC Communication**: Inter-module communication via gRPC with auto-generated code
- **Dependency Injection**: Uber's Fx for clean dependency management
- **Configuration**: Viper-based YAML configuration
- **Structured Logging**: Logrus with multiple formatters (pretty, JSON, text)
- **Database**: GORM with PostgreSQL (SQLite available for testing)
- **Input Validation**: Request validation using go-playground/validator
- **Testing**: Mock generation and coverage reporting
- **CI/CD**: GitHub Actions workflow for automated testing
- **Docker**: Multi-stage builds and docker-compose setup

## Contributing

Contributions are welcome! Here's how you can help:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please ensure your code:
- Passes all tests (`task test`)
- Follows existing code style
- Includes tests for new functionality

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
