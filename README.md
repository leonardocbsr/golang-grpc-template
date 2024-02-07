<img src="https://go.dev/images/go-logo-blue.svg" width="150" height="150">

[![Go](https://github.com/LeoCourbassier/golang-grpc-template/actions/workflows/go.yml/badge.svg)](https://github.com/LeoCourbassier/golang-grpc-template/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/LeoCourbassier/golang-grpc-template)](https://goreportcard.com/report/github.com/LeoCourbassier/golang-grpc-template)
[![codecov](https://codecov.io/gh/LeoCourbassier/golang-grpc-template/branch/main/graph/badge.svg?token=3Z3Z3Z3Z3Z)](https://codecov.io/gh/LeoCourbassier/golang-grpc-template)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
# Go gRPC Template

> [!NOTE]
> The project is still a work in progress, and there are many features that are not yet implemented. 
> 
> Also please note that the gRPC clients are not using any credentials. This is not secure, and should not be used in production. 
> 
> Please feel free to contribute to the project.
> 
This is a template for gRPC modular monoliths in Go. It is designed to be a starting point for a new project, with all the boilerplate code already set up.

We'll use gRPC as the communication protocol between the different modules, and we'll use the `protoc` compiler to generate the necessary code.


## Project Structure

The project is divided into directories:
1. `modules`: This directory contains the different modules of the application. Each module is a self-contained package that contains its own gRPC service implementation. Each module is basically a domain. It contains:
   1. Controllers (gRPC Layer)
   2. Application (Service Layer)
   3. Exceptions (Per module exceptions)
   4. Models (Domain Layer)
   5. Repositories (Data Layer)
   6. Views (gRPC Layer)
2. `proto`: This directory contains the `.proto` files that define the gRPC services and messages. They are separated into different folders for each module.
3. `common`, `config`, `db`, `logger`: These directories contain the common code that is used by all the modules. They are not specific to any module, and are used to provide common functionality to all the modules.

## Packages Used
1. `github.com/golang/protobuf`: This package is used to define the `.proto` files and generate the gRPC code.
2. `google.golang.org/grpc`: This package is used to create the gRPC server and client.
3. `github.com/jinzhu/gorm`: This package is used to interact with the database.
4. `github.com/sirupsen/logrus`: This package is used for logging.
5. `github.com/spf13/viper`: This package is used for configuration management.
6. `github.com/uber-go/fx`: This package is used for dependency injection.
7. `github.com/go-playground/validator`: This package is used for input validation.

## How to Use
1. Clone the repository
2. Run `go mod tidy` to install the dependencies
3. Run `make proto` to generate the gRPC code
4. Run `make run` to start the server
5. ~~Run `make test` to run the tests~~
   
## Features
1. Modular Monolith: The project is divided into different modules, each of which is a self-contained package. This makes it easy to reason about the code, and makes it easy to scale the project.
2. gRPC: The communication between the different modules is done using gRPC. This makes it easy to define the service and messages, and the code is generated automatically.
3. Dependency Injection: We use the `uber-go/fx` package for dependency injection. This makes it easy to manage the dependencies between the different modules.
4. Configuration Management: We use the `spf13/viper` package for configuration management. This makes it easy to manage the configuration of the application.
5. Logging: We use the `sirupsen/logrus` package for logging. This makes it easy to log the different events in the application.
6. Database: We use the `jinzhu/gorm` package to interact with the database. This makes it easy to interact with the database, and provides a lot of functionality out of the box.
7. Input Validation: We use the `go-playground/validator` package for input validation. This makes it easy to validate the input to the different services.
8. We provide injectable clients for the different modules, so that they can communicate with each other.
9. We also provide a `config.yaml` file to manage the configuration of the application including the gRPC server port and the database connection string.


## What's Next
- [ ] Add more error handling
- [ ] Add more tests
- [ ] Add more documentation