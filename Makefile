up:
	docker-compose build && docker-compose up
test:
	go test -cover -race ./...

vegeta-attack:
	vegeta attack -duration=120s -rate=100 -targets=./tests/vegeta-target.list -output=./tests/attack-5.bin

vegeta-report:
	vegeta plot -title=Attack%20Results ./tests/attack-5.bin > ./tests/results.html

