cli:
	go build -o cmarket-cli cmd/cli/main.go && ./cmarket-cli

server:
	go build -o cmarket-server cmd/server/main.go && ./cmarket-server

test:
	go test ./...

lint:
	golangci-lint run ./...
