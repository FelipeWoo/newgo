# Template Go Project

## Overview

This is a general-purpose Go project template designed for rapid development and clean structure.
It includes environment configuration, structured logging, HTTP endpoints, testing, helper scripts, and a Makefile with common automation tasks.

## Features

- Structured code with `cmd/` and `internal/`
- Environment loading from `.env`
- Structured logging with `zerolog`
- Minimal HTTP API using the Go standard library
- Test suite with `go test`
- Formatting and static checks with `gofmt` and `go vet`
- Makefile with common automation tasks

## Requirements

- Go 1.24 or higher
- Unix-like system or WSL (recommended)

## Installation

Clone the repository and run:

```bash
make init
```

This will download module dependencies and prepare the local workspace.

## Running the Project

```bash
make run
```

The server starts on `http://localhost:8000`.

## Environment Configuration

Create a `.env` file and customize values as needed.

```dotenv
ENV=development
LOG_LEVEL=DEBUG
PORT=8000
```

## Testing

```bash
make test
make cov
```

## Linting and Formatting

```bash
make lint
make format
make check
```

## Project Structure

```text
project/
├── cmd/api/            # Application entrypoint
├── internal/           # Private application code
├── tests/              # Integration or smoke tests
├── scripts/            # Helper and scaffolding scripts
├── data/               # Input or reference data
├── logs/               # Application logs
├── docs/               # Documentation
├── Makefile            # Command shortcuts
├── .env.example        # Sample environment variables
├── go.mod              # Go module definition
└── README.md           # Project information
```

## License

This project is licensed under the terms of the LICENSE file.
