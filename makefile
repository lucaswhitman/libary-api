TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

test:
	go test ./...
build:
	go build -ldflags "-X main.version=$(TAG)"
build-linux:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o library-api && docker build -t library-api-amd64:v1 .
pack: build
	docker build -t github.com/lucaswhitman/library-api:$(TAG) .
upload:
	docker push github.com/lucaswhitman/library-api:$(TAG)
local:
	docker-compose build && docker-compose up -d