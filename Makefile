.PHONY: build clean

build:
	mkdir bin
	go build -o bin/open-dash

clean:
	rm -r bin/