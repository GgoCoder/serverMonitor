GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
mkdir release

.PHONY: 
all: clean

.PHONY: 
clean:
	rm -rf release
	$(GOCLEAN)

.PHONY: 
ms: 
	go build -v -a -o ./release/bin/monitorService main.go

.PHONY: 
build:
	
