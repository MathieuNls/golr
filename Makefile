test:
	go test -race -coverprofile=coverage.txt -covermode=atomic

build:
	go get -t -v ./...