# Gosearch

A simple search microservice created using golang, elasticsearch, gRPC and Docker. 

This was created as a learning activity, but if there's any interest in this I'll try to make it into a full open source project

## Usage

Prerequisites: golang, protoc, docker

To setup the project and start the server, first clone this repository and run the following

```
# Compile the protobuf files
./build/proto-gen.sh 

# Get dependencies for the service, including the elasticsearch image, and start the server
docker-compose up --build
```

To test if it's working

```
go run ./cmd/gRPC/client/main.go
```