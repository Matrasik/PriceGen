include .env

BINARY_NAME=main.out
PROJECT_NAME=$(shell basename "$(PWD)")

STDERR=/tmp/.$(PROJECTNAME)-stderr.txt
MAKE_FLAGS += --silent

build:
	go build -o ${BINARY_NAME} main.go

run:
	docker-compose up -d

clean:
	go clean
	rm ${BINARY_NAME}
test:
	go test ./config/ -cover
