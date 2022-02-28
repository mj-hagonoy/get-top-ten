.PHONY: clean build run

clean:
	rm -rf ./build go.sum
	go mod tidy

build: clean
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o build/main main.go

run: build
	./build/main

build-win: clean
	go build -ldflags="-s -w" -o build/main.exe main.go

run-win: build-win
	./build/main.exe
