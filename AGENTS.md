# AGENTS.md - hostname

## Overview
- **Language**: Go (Golang).
- **Purpose**: A simple webapp that prints the hostname and environment variables.
- **Framework**: `gorilla/mux` for routing.
- **Port**: `8080`.

---

## Key Commands
### Run Locally
```shell
go run main.go
```

### Build and Run with Docker
```shell
# Build
docker build -t dirc/hostname:latest .
docker build -t dirc/hostname:$TAG .

# Run
docker run -d --rm -p 8080:8080 dirc/hostname:latest

# Push
docker push dirc/hostname:latest
docker push dirc/hostname:$TAG
```

### Deploy to Kubernetes
```shell
kubectl create deployment web --image=dirc/hostname --replicas=3
kubectl expose deployment web --port 8080
kubectl port-forward svc/web 8080
```

### Verify
```shell
curl http://localhost:8080
```

---

## Code Structure
- **Entry Point**: `main.go`.
- **Endpoints**:
  - `/`: Returns the hostname.
  - `/env`: Returns the value of the `FOO` environment variable (or a message if not set).
- **Dependencies**:
  - `github.com/gorilla/mux` v1.8.1.

---

## Constraints
- No linting, formatting, or testing frameworks are configured.
- No CI/CD workflows are defined.
- No monorepo structure; single Go module (`hostname`).
- No pre-commit hooks or task runners (e.g., `Makefile`, `Justfile`).

---

## Agent-Specific Notes
- **No implicit build steps**: The codebase is simple and does not require codegen, migrations, or complex build steps.
- **No testing**: There are no tests or testing frameworks configured. If tests are added, they should be placed in a `*_test.go` file.
- **No environment variables**: The only environment variable used is `FOO` for the `/env` endpoint.
- **No framework quirks**: The codebase uses standard Go and `gorilla/mux` without customizations.