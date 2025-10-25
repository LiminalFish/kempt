build:
	go build -o bin/app

run: build
	./bin/app

build-front:
	go build cmd/WebFrontRouter -o bin/front

run-front: build-front
	./bin/front
