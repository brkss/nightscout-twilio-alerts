APP_NAME=ns-alerts
SRC_DIR=src

.PHONY: all build clean run

all: build

build:
	go build -o $(APP_NAME) $(SRC_DIR)/main.go

run: build
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)