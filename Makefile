.PHONY: build clean

build:
	mkdir bin
	go build -o bin/open-dash ./src

clean:
	rm -r bin/