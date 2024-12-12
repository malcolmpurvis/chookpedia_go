.PHONY: build clean bee

build:
	go build -o db-management ./cmd/populator/main.go
	go build -o server server.go

clean:
	rm server db-management

bee:
	bee run -main=server.go server

checkstyle:
	gofmt -s -w .
	go vet ./...
	golangci-lint run --enable-all --disable lll,gomnd,exportloopref,execinquery
