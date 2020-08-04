dev-up:
	docker-compose build && docker-compose up

install:
	brew install golang-migrate &
	go get -u github.com/onsi/ginkgo/ginkgo &
	go get -u github.com/onsi/gomega/... &
	go get -u github.com/golang/mock/mockgen

test:
	go test -cover --race ./...