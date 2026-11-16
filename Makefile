# Makefile created following https://makefiletutorial.com/ adapted to this go project. 
# Please use this as a guideline for further implementation.

TARGET := ./cmd/instead/main.go
BUILD_DIR := ./build/instead

run:
	go run $(TARGET)

build:
	go build -o $(BUILD_DIR) $(TARGET)
