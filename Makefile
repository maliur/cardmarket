cli:
	go build -o cmarket-cli cmd/cli/main.go && ./cmarket-cli

server:
	go build -o cmarket-server cmd/server/main.go && ./cmarket-server

test:
	go test ./...

lint:
	golangci-lint run ./...

.PHONY: compile
compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o cmarket-linux_amd64 ./cmd/server/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -installsuffix cgo -o cmarket-windows_amd64.exe ./cmd/server/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -a -installsuffix cgo -o cmarket-linux_arm ./cmd/server/main.go

.PHONY: clean
clean:
	rm -f cmarket-linux_amd64 cmarket-windows_amd64.exe cmarket-linux_arm