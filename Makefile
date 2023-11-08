.PHONY: build
build:
		go build -o main cmd/rayan-gosling/main.go

.PHONY: run
run:
		go run cmd/rayan-gosling/main.go