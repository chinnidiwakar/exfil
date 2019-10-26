# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

VERSION=$(shell grep -e 'VERSION = ".*"' version.go | cut -d= -f2 | sed  s/[[:space:]]*\"//g)
BUILD_DIR=build
NAME=exfil

.PHONY: help, test, clean, deps, versions

all: build

build:
	mkdir -p $(BUILD_DIR)
	cd cmd; $(GOBUILD) -o ../$(BUILD_DIR)/$(NAME) -v
	cd cmd; env GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ../$(BUILD_DIR)/$(NAME)-$(VERSION)-darwin-amd64 -v
	cd cmd; env GOOS=linux GOARCH=amd64 $(GOBUILD) -o ../$(BUILD_DIR)/$(NAME)-$(VERSION)-linux-amd64 -v
	cd cmd; env GOOS=windows GOARCH=amd64 $(GOBUILD) -o ../$(BUILD_DIR)/$(NAME)-$(VERSION)-windows-amd64.exe -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(NAME)
run:
	cd cmd; $(GOBUILD) -o ../$(NAME) -v
	./$(NAME)
deps:
	$(GOGET) ./...
	$(GOGET) github.com/rangertaha/exfil
