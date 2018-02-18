TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

test:
	go test ./...
build:
	go build -ldflags "-X main.version=$(TAG)"
pack: build
	docker build -t github.com/lucaswhitman/library-api:$(TAG) .
upload:
	docker push github.com/lucaswhitman/library-api:$(TAG)
local:
	docker-compose build && docker-compose up -d