all: test build

build: deps
	go build -o api main.go

docker: docker/build docker/push docker/run

docker/build:
	docker build -t joshrendek/the-counter:latest .

docker/push:
	docker push joshrendek/the-counter:latest

docker/run:
	docker run -p 8080:8080 joshrendek/the-counter

docker/test:
	docker build -t joshrendek/the-counter-test -f Dockerfile.test .
	docker run joshrendek/the-counter-test

test: test/unit

test/unit:
	go test -v $(glide nv) -cover

dev:
	go run main.go

dev/race:
	go run -race main.go

lint:
	gometalinter ./... --disable=gotype --disable=gocyclo

linter/install:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

deps:
	glide install
