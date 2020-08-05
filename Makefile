up:
	docker-compose build && docker-compose up
test:
	go test -cover -race ./...