
# Golang webapp - hostname

A simple webapp that:

- prints the hostname
- can also show an environment variable

## Run

```shell
# Kubernetes
kubectl create deployment web --image=dirc/hostname --replicas=3
kubectl expose deployment web --port 8080
kubectl port-forward svc/web 8080

# Docker
docker run -d --rm -p 8080:8080 dirc/hostname:latest
docker run -d --rm -e FOO=BAR -p 8080:8080 dirc/hostname:latest

# Verify
curl http://localhost:8080
curl http://localhost:8080/env

```

## Docker

```shell
TAG=1.4
IMAGE=dirc/hostname

docker build -t ${IMAGE}:$TAG .
docker build -t ${IMAGE}:latest .

docker run -d --rm -p 8080:8080 ${IMAGE}:$TAG

docker push ${IMAGE}:$TAG
docker push ${IMAGE}:latest

```

## webapp

Run the Go webapp:

```shell
go run main.go
```

## Linting

Install and run `golangci-lint` locally:

```shell
# Install golangci-lint (compatible with Go 1.26)
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.0

# Run linting
golangci-lint run
```