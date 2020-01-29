DOCKER_TAG=jcoquintas/health:v0.0.1
DOCKER_RUN=docker run -v "$(PWD):/go/src/health" -w /go/src/ping/service --rm

.PHONY: build test clean image publish

all: lint test build

test:
	@$(DOCKER_RUN) golang go test

lint:
	@$(DOCKER_RUN) golang sh -c "go get golang.org/x/lint/golint && golint"

build:
	@$(DOCKER_RUN) golang go build

image:
	@docker build -t $(DOCKER_TAG) .

publish: image
	@docker push $(DOCKER_TAG)

run-local:
	docker-compose up --build

smoke-test:
	@echo 'Running the smoke tests...'

deploy:
	@echo 'Deploying..'

