APP_NAME=truco
BIN_DIR=bin

run: build
	./$(BIN_DIR)/$(APP_NAME)
	
.PHONY: test
test:
	go test -v -cover ./...

.PHONY: build
build:
	mkdir -p bin
	go build -o ./$(BIN_DIR)/$(APP_NAME) .
