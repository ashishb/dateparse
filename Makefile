build:
	GO111MODULE=on go build -v -o bin/dateparse src/*.go

go_lint:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go vet ./src
	golint -set_exit_status ./src/...
	go tool fix src/
	golangci-lint run

lint: format go_lint

format:
	go fmt ./src/...

clean:
	GO111MODULE=on go clean --modcache
	rm -rf bin/*

test:
	GO111MODULE=on go test ./src/... -v
