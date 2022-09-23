GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
#mkdir ./release

.PHONY: 
all: clean

.PHONY: 
clean:
	rm -rf release
	$(GOCLEAN)

.PHONY: 
ms: 
	go build -v -a -o ./release/bin/monitorService ./cmd/webService/main.go

.PHONY: 
build:
	
.PHONY:
grpc: logGrpc userGrpc

.PHONY:
logGrpc:
	protoc --go_out=./internal/logService/proto/ ./internal/logService/proto/log.proto
	protoc --go-grpc_out=./internal/logService/proto/ ./internal/logService/proto/log.proto
	@echo "genetated log grpc proto successfully!\n"

.PHONY:
userGrpc:
	protoc  --go_out=./internal/userService/proto/ ./internal/userService/proto/user.proto
	protoc  --go-grpc_out=./internal/userService/proto/ ./internal/userService/proto/user.proto
	@echo "genetated user grpc proto successfully!\n"




.PHONY:
test:
	go test -v -cover ./...
