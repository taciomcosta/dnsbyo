BINARY_NAME=dnsbyo

run: build
	./$(BINARY_NAME)
build:
	go build -o $(BINARY_NAME)
clean:
	rm $(BINARY_NAME)
