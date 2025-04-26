.PHONY: build run clean

build:
	go build -o bin/rate-limiter cmd/server/main.go

run: build
	./bin/rate-limiter

clean:
	rm -rf bin/