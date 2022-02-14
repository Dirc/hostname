
# Golang webapp - hostname

A simple webapp that prints the hostname

## Run

```shell
# Kubernetes
kubectl create deployment web --image=dirc/hostname --replicas=3
kubectl expose deployment web --port 8080
kubectl port-forward svc/web 8080

# Docker
docker run -d --rm -p 8080:8080 dirc/hostname:1.0

# Verify
curl http://localhost:8080

```

## Docker

```shell
TAG=1.1
IMAGE=dirc/hostname

docker build -t $IMAGE:$TAG .
docker build -t $IMAGE:latest .

docker run -d --rm -p 8080:8080 $IMAGE:$TAG

docker push $IMAGE:$TAG
docker push $IMAGE:latest

```

## webapp

Run go webapp

```shell
go run app.go

```