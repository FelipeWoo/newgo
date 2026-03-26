# newgo

`newgo` is a lightweight Go starter template plus a small scaffold script for creating new repositories from it.

It includes:

- A minimal HTTP API built with the Go standard library
- Structured logging with `zerolog`
- Environment-aware config loading
- `Makefile` tasks for common development flows
- Shell scripts to clone the template and install a `newgo` command

## Requirements

- Go 1.24.2 or newer
- Git
- Bash
- `curl` if you want to install the global `newgo` helper

## What This Repo Contains

### 1. The template app

The generated app starts an HTTP server from [`cmd/api/main.go`](/home/zenon/Documents/300_Projects/SOFTWARE/PUBLIC/newgo/cmd/api/main.go) and registers basic routes from [`internal/router/routes/system.go`](/home/zenon/Documents/300_Projects/SOFTWARE/PUBLIC/newgo/internal/router/routes/system.go):

- `GET /` returns `{"status":"ok"}`
- `GET /health` returns `{"service":"running"}`

### 2. The repo generator

The script [`scripts/make_repo_newgo.sh`](/home/zenon/Documents/300_Projects/SOFTWARE/PUBLIC/newgo/scripts/make_repo_newgo.sh) clones this template into a new folder, removes template artifacts, initializes a fresh git repo, creates a starter `.env`, and adds a private notes folder.

The helper installer [`scripts/alias_make_repo_newgo.sh`](/home/zenon/Documents/300_Projects/SOFTWARE/PUBLIC/newgo/scripts/alias_make_repo_newgo.sh) downloads that generator script to your home directory and exposes it as a shell function named `newgo`.

## Create A New Project

### Option 1: run the generator directly

```bash
bash scripts/make_repo_newgo.sh my-api
```

### Option 2: install the `newgo` shell helper

```bash
bash scripts/alias_make_repo_newgo.sh
source ~/.bash_aliases
newgo my-api
```

After generation:

```bash
cd my-api
make init
make run
```

## Local Development

From this repository:

```bash
make init
make run
```

Useful commands:

- `make test` runs the test suite
- `make cov` runs coverage and prints the summary
- `make lint` runs `go vet`
- `make format` runs `gofmt`
- `make check` runs format, lint, and tests
- `make clean` removes local build/test artifacts
- `make reset` recreates local artifacts from scratch

## Configuration

Runtime config is loaded by [`internal/config/config.go`](/home/zenon/Documents/300_Projects/SOFTWARE/PUBLIC/newgo/internal/config/config.go).

Supported variables:

- `APP_NAME`
- `LOG_LEVEL`
- `PORT`
- `ENV`

Current environment file behavior:

- `make init` reads `APP_NAME` from a root `.env` file to rename the Go module
- Application runtime currently looks for env files under `shared/.env`, `shared/.env.test`, and `shared/.env.prod`

If you are using this template as-is, keep that mismatch in mind when preparing local config.

## Project Layout

```text
.
├── cmd/api/                 # Application entrypoint
├── internal/bootstrap/      # Startup lifecycle
├── internal/config/         # Env/config loading
├── internal/logger/         # Logging setup and helpers
├── internal/router/         # HTTP server and routes
├── scripts/                 # Project generator and helper scripts
├── tests/                   # Basic test setup
├── docs/                    # Documentation notes/templates
├── data/                    # Example raw data
├── Makefile                 # Dev commands
└── README.md
```

## License

This project is licensed under the terms in [`LICENSE`](/home/zenon/Documents/300_Projects/SOFTWARE/PUBLIC/newgo/LICENSE).
