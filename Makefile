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
logGrpc:
	protoc --go_out=./internal/logService/proto/ ./internal/logService/proto/log.proto
	@echo "genetated log grpc successfully!"

.PHONY:
test:
	go test -v -cover ./...
