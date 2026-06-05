FROM golang:1.26-alpine

WORKDIR /app
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY main.go ./
RUN go build -v -o /hostname .

FROM scratch
COPY --from=0 /hostname /hostname
EXPOSE 8080
CMD ["/hostname"]
