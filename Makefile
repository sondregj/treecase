all: fmt vet test build

fmt:
	go fmt ./...

vet:
	go vet ./...

build:
	go build -o ./bin/treecase .

test:
	go test ./...
