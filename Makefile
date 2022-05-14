all: go

exec = fizzbuzz-leboncoin

start:
	@docker-compose up

build:
	@docker-compose build

stop:
	@docker-compose down

restart: stop start

go: main.go
	go test ./app/handlers
	go build -o bin/$(exec) main.go
	go generate

clean:
	rm -f bin/$(exec) *~ *#
	rm -rf src/gopkg.in
	go mod tidy

re: clean all
