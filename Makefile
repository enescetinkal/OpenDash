.PHONY: build clean run

build:
	mkdir -p bin
	go build -o bin/open-dash ./src

clean:
	rm -r bin/

run:
	go run ./src