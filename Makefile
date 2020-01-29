DOCKER_TAG=jcoquintas/health:v0.0.1
DOCKER_RUN=docker-compose run -w /go/src/app --rm web

.PHONY: build test clean image publish

all: lint test build

test:
	@$(DOCKER_RUN) go test

lint:
	@$(DOCKER_RUN) sh -c "go get golang.org/x/lint/golint && golint"

build:
	@$(DOCKER_RUN) go build

image:
	@docker-compose build

publish: image
	@docker-compose push

run-local:
	docker-compose up --build

smoke-test:
	@echo 'Running the smoke tests...'

deploy:
	@echo 'Deploying..'

