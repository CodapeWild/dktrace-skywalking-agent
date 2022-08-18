GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
NAME = $(shell basename $(CURDIR))
SkyApiPath = $(shell go env GOPATH)/src/skywalking.apache.org/repo

default: deps build

deps:
	@echo create directory: $(SkyApiPath)
	@mkdir -p $(SkyApiPath)
	@cd $(SkyApiPath);git clone -v git@github.com:apache/skywalking-goapi.git goapi

build:
	@echo build binary for $(GOOS)/$(GOARCH)
	@go build -o $(NAME)-$(GOOS)-$(GOARCH)