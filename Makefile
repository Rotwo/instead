# Makefile created following https://makefiletutorial.com/ adapted to this go project. 
# Please use this as a guideline for further implementation.

TARGET := ./cmd/instead/main.go
BUILD_DIR := ./build
OUTPUT := $(BUILD_DIR)/instead

run:
	go run $(TARGET)

build:
	go build -o $(OUTPUT) $(TARGET)

.PHONY: clean
clean:
	rm -r $(BUILD_DIR)
