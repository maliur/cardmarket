run:
	go build -o cmarket *.go && ./cmarket

test:
	go test ./...

lint:
	golangci-lint run ./...
