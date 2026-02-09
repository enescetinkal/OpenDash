.PHONY: build clean

build:
	mkdir -p bin
	go build -o bin/open-dash ./src

clean:
	rm -r bin/