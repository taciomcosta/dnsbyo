BINARY_NAME=dnsbyo
MAIN=github.com/taciomcosta/dnsbyo/main

run: build
	./$(BINARY_NAME)
build:
	go build -o $(BINARY_NAME) $(MAIN)
clean:
	rm $(BINARY_NAME)
