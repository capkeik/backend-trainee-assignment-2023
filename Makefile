run:
	chmod +x cmd/segmentation/wait-for-it.sh ; docker-compose up  --remove-orphans --build

lint:
	gofmt -l -w .
	golangci-lint run --fix