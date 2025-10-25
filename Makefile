build:
	go build -o bin/app

run: build
	./bin/app

run-front:
	go run ./cmd/webFrontRouter
