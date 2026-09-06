BIN := bin/mr-papaya

.PHONY: build check test vet run clean

build:
	go build -o $(BIN) ./cmd/mr-papaya

check: build vet test

test:
	go test ./...

vet:
	go vet ./...

run:
	go run ./cmd/mr-papaya

clean:
	rm -rf bin dist
