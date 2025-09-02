# Kratos Project Template

## Go 1.25+ Optimizations

This project uses Go 1.25.0+ with built-in container-aware GOMAXPROCS feature. The `go.uber.org/automaxprocs` package has been removed since Go 1.25+ now:

- Detects container CPU limits and quotas
- Auto adjusts GOMAXPROCS to match container resources
- Provides better performance in containerized environments without external dependencies

See the [Go 1.25 release notes](https://tip.golang.org/doc/go1.25) to learn more.

## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

