export GO111MODULE=on

fmt:
	go fmt ./...

tidy:
	go mod tidy
