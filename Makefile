.PHONY: build
build:
		go build -o main cmd/rayan-gosling/main.go
		go build -o main cmd/tgbot/tgbot.go

.PHONY: run1
run1:
		go run cmd/rayan-gosling/main.go

.PHONY: run2
run2:
		go run cmd/tgbot/tgbot.go