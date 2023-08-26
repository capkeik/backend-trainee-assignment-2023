run:
	docker-compose up  --remove-orphans --build

lint:
	gofmt -l -w .
	golangci-lint run --fix