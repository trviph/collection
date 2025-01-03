install_lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

setup: install_lint

lint:
	golangci-lint run

test: lint
	go test -coverprofile=./test_profile ./... 

cover: test
	go tool cover -html=./test_profile && unlink ./test_profile