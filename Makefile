.PHONY: build clean bee

build:
	go build -o db-management ./cmd/populator/main.go
	go build -o server ./cmd/server/main.go

clean:
	rm server db-management

bee:
	bee run -main=cmd/server/main.go server

checkstyle:
	gofmt -s -w .
	go vet ./...
	golangci-lint run --enable-all --disable lll,gomnd,exportloopref,execinquery
