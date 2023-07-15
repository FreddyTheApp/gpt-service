test:
	go test -cover ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

.PHONY: test cover